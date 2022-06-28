package main

import (
	"log"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/server"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/handler"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/service"
)

func main() {
	// declaring dependencies
	repos := service.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// initializing server instance, and check for error
	srv := new(server.Server)
	if err := srv.RunServer("localhost:8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}
}
