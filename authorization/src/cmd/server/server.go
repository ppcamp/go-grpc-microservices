package main

import (
	"net"
	"os"
	"os/signal"
	"streamer/configs"
	"streamer/controllers/chat"
	"streamer/utils"

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

func serve(grpcServer *grpc.Server, ch chan os.Signal) {
	listener := utils.Must(net.Listen("tcp", *configs.APP_PORT))

	server := chat.NewServer()
	chat.RegisterChatServiceServer(grpcServer, server)

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatal(utils.Wraps("Fail to serve gRPC", err))
		ch <- os.Interrupt
	}
}
