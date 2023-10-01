package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AxzelBC/ToDo-Go/internal/app/port/IServer"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/config"
	"github.com/gin-gonic/gin"
)

// Host address
const defaultHost = "0.0.0.0"

// Struct config of the server and port
type httpServer struct {
	Port   uint
	server *http.Server
}

// create a httpConfig instance
func NewHttpConfig(router *gin.Engine, config config.HttpServerConfig) IServer.IHttpServer {
	return &httpServer{
		Port: config.Port,
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%d", defaultHost, config.Port),
			Handler: router,
		},
	}
}

// Run server
func (s *httpServer) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf(
				"Failed to stater HttpServer listen port %d, err=%s\n",
				s.Port,
				err.Error(),
			)
			return
		}
	}()

	log.Printf("Start Service with port %d\n", s.Port)
	return
}

// Close server
func (s *httpServer) Stop() {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(3)*time.Second,
	)

	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown err=%s\n", err.Error())
		return
	}

	return
}
