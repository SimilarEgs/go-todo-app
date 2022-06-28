package main

import (
	"fmt"
	"log"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/server"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/handler"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/service"
	"github.com/spf13/viper"
)

func main() {

	// read confing file and handle erros
	if err := InitConfig(); err != nil {
		log.Fatalf("[Error] failed to load config file: %s\n", err.Error())
	}

	// declaring dependencies
	repos := service.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	fmt.Println(viper.GetString("port"))

	// initializing server instance, and check for error
	srv := new(server.Server)
	if err := srv.RunServer(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
