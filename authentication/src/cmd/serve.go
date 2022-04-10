package main

import (
	"net"
	"os"
	"streamer/configs"
	"streamer/controllers/auth"
	"streamer/controllers/user_password"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/services"
	"streamer/utils"
	"streamer/utils/jwt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func getHandler() *services.Handler {
	cacheRepository := cache.NewCacheRepository()

	privateKey := utils.Must(jwt.ParseSSHPrivateKey(*configs.JWT_PRIVATE))
	signer := jwt.NewJwt(privateKey)

	db := utils.Must(database.NewStore("", ""))

	return &services.Handler{
		Cache:    cacheRepository,
		Signer:   signer,
		Database: db,
	}
}

func serve(grpcServer *grpc.Server, ch chan os.Signal) {
	listener := utils.Must(net.Listen("tcp", *configs.APP_PORT))

	handler := getHandler()

	authServer := auth.NewAuthService(handler)
	userServer := user_password.NewUserPasswordService(handler)

	auth.RegisterAuthServiceServer(grpcServer, authServer)
	user_password.RegisterUserPasswordServiceServer(grpcServer, userServer)

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Fatal(utils.Wraps("Fail to serve gRPC", err))
		ch <- os.Interrupt
	}
}
