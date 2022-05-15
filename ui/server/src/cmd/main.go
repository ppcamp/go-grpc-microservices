package main

import (
	nhttp "net/http"
	"os"
	"os/signal"
	"syscall"

	"ui/server/src/configs"
	"ui/server/src/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "js-framework-server"
	app.Usage = "Used to serve any js page app"
	app.Flags = configs.Flags
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) error {
	gin.SetMode(gin.DebugMode)
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	ctx, stop := signal.NotifyContext(c.Context, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	errChan := make(chan error, 1)

	s := http.NewServer(configs.ServerPort)
	s.Setup(configs.Path)

	go func() {
		if err := s.Serve(); err != nil && err != nhttp.ErrServerClosed {
			log.Error(err)
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		log.WithError(err).Println("Some error occurred!")
		return err

	case <-ctx.Done():
		log.Info("Shutting down server gracefully, press Ctrl+C again to force it")
		stop()

		if err := s.GracefulShutdown(); err != nil {
			log.WithError(err).Println("Server failed to shutdown")
		}

		log.Info("Server exiting")
		return nil
	}
}
