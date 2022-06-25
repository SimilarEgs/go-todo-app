package main

import (
	"log"

	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/server"
)

func main() {

	// initializing server instance, and check for error
	srv := new(server.Server)
	if err := srv.RunServer("localhost:8080"); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}
}
