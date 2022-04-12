package main

import (
	"fmt"
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
	cacheConfig := cache.CacheConfig{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	}
	cacheId := fmt.Sprintf("%s-%s", configs.APP_NAME, *configs.APP_ID)
	cacheRepository := utils.Must(cache.NewCacheRepository(cacheConfig, cacheId))

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
