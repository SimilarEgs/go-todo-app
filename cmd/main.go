package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/SimilarEgs/go-todo-app/internal/server"
	"github.com/SimilarEgs/go-todo-app/logger"
	"github.com/SimilarEgs/go-todo-app/pkg/handler"
	"github.com/SimilarEgs/go-todo-app/pkg/repository"
	"github.com/SimilarEgs/go-todo-app/pkg/service"
	"github.com/SimilarEgs/go-todo-app/utils"
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

	// initializing server instance
	srv := new(server.Server)

	// starting server
	go func() {
		if err := srv.RunServer(config.ServerPort, handlers.InitRoutes()); err != nil {
			log.Fatalf("[Error] failed to start server: %s", err.Error())
		}
	}()

	log.Println("[Info] TodoApp started")

	// implementing graceful shutdown

	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // blocks untill either SIGINT or SIGTERM is received

	// affter received signal
	// call ShutDown method and close all db connections

	log.Println("[Info] TodoApp shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.ShutDownServer(ctx); err != nil {
		log.Println("[Error] occurred while shutting down the server: %v", err)
	}

	if err := db.Close(); err != nil {
		log.Println("[Error] occurred while closing db connection: %v", err)
	}
}
