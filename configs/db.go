package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Host     string `envconfig:"DB_HOST" default:"localhost"`
	Port     string `envconfig:"DB_PORT" default:"5432"`
	User     string `envconfig:"DB_USER" default:"postgres"`
	Password string `envconfig:"DB_PASSWORD" default:"password"`
	Database string `envconfig:"DB_NAME" default:"postgres"`
}

var DBConfig Database

func LoadConfig() {
	err := envconfig.Process("test-backend", &DBConfig)
	if err != nil {
		log.Fatal("Error loading config: ", err)
	}
}
