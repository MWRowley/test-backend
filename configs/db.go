package configs

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"password"`
	Database string `envconfig:"DB_DATABASE" default:"postgres"`
}

var DBConfig Database

func LoadDBConfig() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("No .env file found")
	}

	err = envconfig.Process("", &DBConfig)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
}
