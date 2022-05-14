package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"authentication/configs"
	"authentication/helpers/handlers"
	"authentication/http/gRPC/auth"
	"authentication/http/gRPC/user_password"
	"authentication/repositories/cache"
	"authentication/repositories/database"
	grpcutils "authentication/utils/grpc"
	"authentication/utils/jwt"

	"github.com/ppcamp/go-lib/errors"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
)

var (
	grpcServer *grpc.Server
	server     *http.Server
)

func inits() {
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
}

func main() {
	app := cli.NewApp()
	app.Name = "microservice-authentication"
	app.Usage = "Used to authorize every request made for user"
	app.Flags = configs.Flags
	app.Action = run
	app.Run(os.Args)
}

func gracefulStop() {
	logrus.Info("Closing server...")
	defer logrus.Info("Server closed!")

	logrus.Info("Closing gRPC connections")
	grpcServer.GracefulStop()

	logrus.Info("Closing http server")
	if err := server.Close(); err != nil {
		logrus.Error(err)
	}

	handler := handlers.GetHandler()

	logrus.Info("Closing PostgreSQL connections")
	if err := handler.Database.Close(); err != nil {
		logrus.Error("fail when closing database")
	}

}

func run(c *cli.Context) (err error) {
	inits()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	grpcServer = grpc.NewServer()
	registerGrpc()

	listener := errors.Must(net.Listen("tcp", configs.AppPort))
	server = grpcutils.NewMuxServer(http.NewServeMux(), grpcServer)

	go func() {
		if err := server.Serve(listener); err != nil {
			logrus.Fatal(errors.Wraps("Fail to serve gRPC", err))
			signalChan <- os.Interrupt
		}
	}()

	logrus.Infof("Server started and listening at http://localhost%s", configs.AppPort)
	<-signalChan
	gracefulStop()
	return
}

func registerGrpc() {
	handler := handlers.GetHandler()

	authServer := auth.NewAuthService(handler)
	userServer := user_password.NewUserPasswordService(handler)
	health := &grpcutils.MyGrpcService{}

	grpc_health_v1.RegisterHealthServer(grpcServer, health)

	auth.RegisterAuthServiceServer(grpcServer, authServer)
	user_password.RegisterUserPasswordServiceServer(grpcServer, userServer)
}
