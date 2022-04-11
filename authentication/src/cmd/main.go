package main

import (
	"net"
	"os"
	"os/signal"
	"streamer/configs"
	"streamer/controllers/base"
	"streamer/utils"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	grpcServer := grpc.NewServer()
	listener := utils.Must(net.Listen("tcp", *configs.APP_PORT))
	initServices(grpcServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			logrus.Fatal(utils.Wraps("Fail to serve gRPC", err))
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

	handler := base.GetHandler()

	logrus.Info("Closing PostgreSQL connections")
	if err := handler.Database.Close(); err != nil {
		logrus.Error("fail when closing database")
	}
}
