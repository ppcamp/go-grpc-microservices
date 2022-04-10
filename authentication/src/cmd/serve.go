package main

import (
	"streamer/configs"
	"streamer/controllers/auth"
	"streamer/controllers/user_password"
	"streamer/repositories/cache"
	"streamer/repositories/database"
	"streamer/services"
	"streamer/utils"
	"streamer/utils/jwt"

	"google.golang.org/grpc"
)

func getHandler() *services.HHandler {
	cacheRepository := cache.NewCacheRepository()

	privateKey := utils.Must(jwt.ParseSSHPrivateKey(*configs.JWT_PRIVATE))
	signer := jwt.NewJwt(privateKey)

	db := utils.Must(database.NewStore(*configs.DATABASE_QUERY))

	h := &services.HHandler{
		Cache:    cacheRepository,
		Signer:   signer,
		Database: db,
	}
	services.InitHandler(h)
	return h
}

func init_services(grpcServer *grpc.Server) {
	handler := getHandler()

	authServer := auth.NewAuthService(handler)
	userServer := user_password.NewUserPasswordService(handler)

	auth.RegisterAuthServiceServer(grpcServer, authServer)
	user_password.RegisterUserPasswordServiceServer(grpcServer, userServer)
}
