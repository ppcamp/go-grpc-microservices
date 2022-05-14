package grpc

import (
	"net/http"
	"strings"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func HttpAndGrpcMux(httpHandler http.Handler, grpcHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.HasPrefix(r.Header.Get("content-type"), "application/grpc") {
			grpcHandler.ServeHTTP(w, r)
			return
		}
		httpHandler.ServeHTTP(w, r)
	})
}

func NewMuxServer(httpHandler http.Handler, grpcHandler http.Handler) *http.Server {
	muxHandlers := HttpAndGrpcMux(httpHandler, grpcHandler)
	http2Server := new(http2.Server)

	return &http.Server{Handler: h2c.NewHandler(muxHandlers, http2Server)}
}
