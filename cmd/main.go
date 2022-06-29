package main

import (
	"fmt"
	"log"
	"os"

	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/server"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/handler"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/repository"
	"github.com/SimilarEgs/CRUD-TODO-LIST/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {
	// read confing file and handle erros
	if err := InitConfig(); err != nil {
		log.Fatalf("[Error] failed to load config file: %s\n", err.Error())
	}

	// load .env files
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("[Error] .env file didn't load: %s", err.Error())
	}

	//initializing db
	db, err := repository.CreatePostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("db_password"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatalf("[Error] faild to initilize the data base: %s", err.Error())
	}

	// declaring dependencies
	repos := service.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	fmt.Println(viper.GetString("port"))

	// initializing server instance, and check for error
	srv := new(server.Server)
	if err := srv.RunServer(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("[Error] failed to start server: %s", err.Error())
	}
}

// this function will load config.yml file
func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
