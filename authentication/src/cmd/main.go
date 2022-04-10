package main

import (
	"os"
	"os/signal"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	grpcServer := grpc.NewServer()
	go serve(grpcServer, signalChan)

	logrus.Info("Server started and listening")
	<-signalChan
	logrus.Info("Closing server")
	grpcServer.GracefulStop()
	logrus.Info("Server closed")
}
