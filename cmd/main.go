package main

import (
	"log"

	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/handler"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/server"
)

func main() {

	// creaing instance of handler object
	handlers := new(handler.Hanlder)

	// initializing server instance, and check for error
	srv := new(server.Server)
	if err := srv.RunServer("localhost:8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}
}
