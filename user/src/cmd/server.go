package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"streamer/configs"
	"streamer/helpers/handlers"
	"streamer/http/gRPC/user"
	"streamer/microservices/auth"
	"streamer/repositories/cache"
	"streamer/repositories/database"
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
	userService := user.NewUserService(handler)
	user.RegisterUserServiceServer(grpcServer, userService)
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

	logrus.Info("Starting a new store")
	db := errors.Must(database.NewStore(*configs.DATABASE_QUERY))

	logrus.Info("Starting a new user Microservice connection")
	authMs := errors.Must(auth.NewUserPasswordClient())

	logrus.Info("Initializing handlers")
	h := &handlers.Handler{
		Cache:    cacheRepository,
		Database: db,
		Auth:     authMs,
	}
	handlers.Init(h)
	return h
}
