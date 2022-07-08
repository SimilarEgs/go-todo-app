package main

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/server"
	"github.com/SimilarEgs/CRUD-TODO-LIST/logger"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/handler"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/repository"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/service"
	"github.com/SimilarEgs/CRUD-TODO-LIST/utils"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func main() {
	// setting log fortmat to JSON
	log.SetFormatter(&log.JSONFormatter{})

	// read confing file and handle erros
	config, err := utils.InitConfig(".")
	if err != nil {
		log.Fatalf("[Error] failed to load config file: %s", err.Error())
	}

	// load logger
	logger.InitializeLogging()

	//initializing db
	db, err := repository.CreatePostgresDB(repository.Config{
		Host:     config.DBHost,
		Port:     config.DBPort,
		Username: config.DBUserName,
		DBName:   config.DBName,
		Password: config.DBPassword,
		SSLMode:  config.DBSSLMode,
	})
	if err != nil {
		log.Fatalf("[Error] faild to initilize the data base: %s", err.Error())
	}

	// declaring dependencies
	repos := service.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// initializing server instance, and check for error
	srv := new(server.Server)
	if err := srv.RunServer(config.ServerPort, handlers.InitRoutes()); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}
}
