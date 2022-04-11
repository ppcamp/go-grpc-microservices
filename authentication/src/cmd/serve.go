package main

import (
	"streamer/configs"
	"streamer/controllers/auth"
	"streamer/controllers/base"
	"streamer/controllers/user_password"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/utils"
	"streamer/utils/jwt"

	"google.golang.org/grpc"
)

func initServices(grpcServer *grpc.Server) {
	handler := initAndGetHandler()

	authServer := auth.NewAuthService(handler)
	userServer := user_password.NewUserPasswordService(handler)

	auth.RegisterAuthServiceServer(grpcServer, authServer)
	user_password.RegisterUserPasswordServiceServer(grpcServer, userServer)
}

func initAndGetHandler() *base.Handler {
	cacheRepository := cache.NewCacheRepository()

	privateKey := utils.Must(jwt.ParseSSHPrivateKey(*configs.JWT_PRIVATE))
	signer := jwt.NewJwt(privateKey)

	db := utils.Must(database.NewStore(*configs.DATABASE_QUERY))

	h := &base.Handler{
		Cache:    cacheRepository,
		Signer:   signer,
		Database: db,
	}
	base.InitHandler(h)
	return h
}
