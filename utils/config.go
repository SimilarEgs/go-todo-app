package utils

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TokenSecurityKey string `mapstructure:"JWT_SECURITY_KEY"`
	TokenDuration    string `mapstructure:"JWT_TOKEN_DURATION_IN_MINUTE"`
	DBPassword       string `mapstructure:"DB_PASSWORD"`
	DBPort           string `mapstructure:"DB_PORT"`
	DBHost           string `mapstructure:"DB_HOST"`
	DBName           string `mapstructure:"DB_NAME"`
	DBUserName       string `mapstructure:"DB_USERNAME"`
	DBSSLMode        string `mapstructure:"DB_SSLMODE"`
	ServerPort       string `mapstructure:"SERVER_PORT"`
}

// this function load conf file
func InitConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	if err = godotenv.Load("config.env"); err != nil {
		log.Fatalf("[Error] .env file didn't load: %s", err.Error())
	}

	err = viper.Unmarshal(&config)

	return
}
