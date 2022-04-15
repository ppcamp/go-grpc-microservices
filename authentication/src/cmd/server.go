package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"streamer/configs"
	"streamer/helpers/handlers"
	"streamer/http/gRPC/auth"
	"streamer/http/gRPC/user_password"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/utils/jwt"
	"time"

	"github.com/ppcamp/go-lib/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	grpcServer := grpc.NewServer()
	listener := errors.Must(net.Listen("tcp", *configs.APP_PORT))
	initServices(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Fatal(errors.Wraps("Fail to serve gRPC", err))
			signalChan <- os.Interrupt
		}
	}()

	logrus.Info("Server started and listening")
	<-signalChan
	gracefulStop(grpcServer)
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
		Addr:        "localhost:6379",
		Password:    "", // no password set
		DB:          0,  // use default DB
		DialTimeout: time.Second * 2,
	}
	cacheId := fmt.Sprintf("%s-%s", configs.APP_NAME, *configs.APP_ID)
	cacheRepository := errors.Must(cache.NewCacheRepository(cacheConfig, cacheId))

	logrus.Info("Creating vault manager/signer")
	privateKey := errors.Must(jwt.ParseSSHPrivateKey(*configs.JWT_PRIVATE))
	jwt.Init(privateKey)

	logrus.Info("Starting a new store")
	db := errors.Must(database.NewStore(*configs.DATABASE_QUERY))

	logrus.Info("Initializing handlers")
	h := &handlers.Handler{
		Cache:    cacheRepository,
		Database: db,
	}
	handlers.Init(h)
	return h
}
