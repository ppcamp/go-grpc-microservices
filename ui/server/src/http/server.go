package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const CLOSE_TIMEOUT = 3 * time.Second

func (s *server) Setup(staticPath string) {
	s.RedirectTrailingSlash = true

	s.Use(LoggingMiddleware())
	s.Use(requestid.New())
	s.Use(gin.Recovery())
	s.Use(cors.Default())
	s.Use(gzip.Gzip(gzip.DefaultCompression))

	s.Use(SinglePageAppMiddleware("/", staticPath))
	s.GET("/healthcheck", func(ctx *gin.Context) { ctx.String(200, "Working") })
}

type server struct {
	srv *http.Server
	*gin.Engine
}

type Server interface{}

func NewServer(port string) *server {
	srv := &http.Server{Addr: port}
	return &server{srv, gin.New()}
}

func (s *server) Serve() error {
	s.srv.Handler = s
	log.Infof("Starting server at http://localhost%s", s.srv.Addr)
	return s.srv.ListenAndServe()
}

func (s *server) GracefulShutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), CLOSE_TIMEOUT)
	defer cancel()

	err := s.srv.Shutdown(ctx)
	return err
}
