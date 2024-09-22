package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	Port string `envconfig:"SERVER_PORT" default:"3000"`
	Host string `envconfig:"SERVER_HOST" default:"localhost"`
}

var ServerConfig Server

func LoadServerConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found")
	}

	err = envconfig.Process("", &ServerConfig)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
}
