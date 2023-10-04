package cmd

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/config"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/repository"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/middleware"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/rest/routes"
	"github.com/AxzelBC/ToDo-Go/internal/infrastructure/server"

	"github.com/gin-gonic/gin"
)

func Runner() {
	instance := gin.New()

	// add middlewares
	instance.Use(
		gin.Logger(),
		gin.Recovery(),
		middleware.CORSMiddleware(),
	)

	// server config
	httpServer := server.NewHttpConfig(
		instance,
		config.HttpServerConfig{
			Port: 8080,
		},
	)

	// Create DB
	db, err := repository.NewDB(
		&config.DatabaseConfig{},
	)

	if err != nil {
		log.Fatalf("failed to new database err=%s\n", err.Error())
	}

	// Init routes and controllers
	routes.AppRoutes(instance, &db)

	// start the http server
	httpServer.Start()
	defer httpServer.Stop()
	log.Println("listening signals...")

	/* signal of exit */
	c := make(chan os.Signal, 1)
	signal.Notify(
		c,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	<-c
	log.Println("graceful shutdown...")
}
