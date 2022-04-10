package main

import (
	"net"
	"os"
	"streamer/configs"
	"streamer/utils"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func serve(grpcServer *grpc.Server, ch chan os.Signal) {
	listener := utils.Must(net.Listen("tcp", *configs.APP_PORT))

	// server := auth.NewAuthService()
	// auth.RegisterAuthServiceServer(grpcServer, server)

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatal(utils.Wraps("Fail to serve gRPC", err))
		ch <- os.Interrupt
	}
}
