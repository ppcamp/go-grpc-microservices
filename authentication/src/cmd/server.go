package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"

	"authentication/configs"
	"authentication/helpers/handlers"
	"authentication/http/gRPC/auth"
	"authentication/http/gRPC/user_password"
	"authentication/repositories/cache"
	"authentication/repositories/database"
	"authentication/utils/jwt"

	"github.com/ppcamp/go-lib/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := cli.NewApp()
	app.Name = "microservice-authentication"
	app.Usage = "Used to authorize every request made for user"
	app.Flags = configs.Flags
	app.Action = run
	app.Run(os.Args)
}

func run(c *cli.Context) (err error) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	grpcServer := grpc.NewServer()
	listener := errors.Must(net.Listen("tcp", configs.AppPort))
	initServices(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Fatal(errors.Wraps("Fail to serve gRPC", err))
			signalChan <- os.Interrupt
		}
	}()

	logrus.Infof("Server started and listening at http://localhost%s", configs.AppPort)
	<-signalChan
	gracefulStop(grpcServer)
	return
}

func gracefulStop(grpcServer *grpc.Server) {
	logrus.Info("Closing server...")
	defer logrus.Info("Server closed!")

	logrus.Info("Closing gRPC connections")
	grpcServer.GracefulStop()

	handler := handlers.GetHandler()

	logrus.Info("Closing PostgreSQL connections")
	if err := handler.Database.Close(); err != nil {
		logrus.Error("fail when closing database")
	}
}

func initServices(grpcServer *grpc.Server) {
	handler := initAndGetHandler()

	authServer := auth.NewAuthService(handler)
	userServer := user_password.NewUserPasswordService(handler)

	auth.RegisterAuthServiceServer(grpcServer, authServer)
	user_password.RegisterUserPasswordServiceServer(grpcServer, userServer)
}

func initAndGetHandler() *handlers.Handler {
	logrus.Info("Starting connection with cache")
	cacheConfig := cache.CacheConfig{
		Addr:        configs.CacheHost + ":" + configs.CachePort,
		Password:    "",              // no password set
		DB:          configs.CacheDb, // use default DB
		DialTimeout: time.Second * 2,
	}
	cacheId := fmt.Sprintf("%s-%s", configs.APP_NAME, configs.AppId)
	cacheRepository := errors.Must(cache.NewCacheRepository(cacheConfig, cacheId))

	logrus.Info("Creating vault manager/signer")
	privateKey := errors.Must(jwt.ParseSSHPrivateKey(configs.JwtPrivate))
	jwt.Init(privateKey)

	logrus.Info("Starting a new store")

	connQuery := fmt.Sprintf(
		database.CONNECTION_QUERY,
		configs.DatabaseHost,
		configs.DatabasePort,
		configs.DatabaseUser,
		configs.DatabasePassword,
		configs.DatabaseName,
	)
	logrus.Info(connQuery)
	db := errors.Must(database.NewStore(connQuery))

	logrus.Info("Initializing handlers")
	h := &handlers.Handler{
		Cache:    cacheRepository,
		Database: db,
	}
	handlers.Init(h)
	return h
}
