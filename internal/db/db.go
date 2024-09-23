package db

import (
	"database/sql"
	"fmt"
	"log"
	"test-backend/configs"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

var DB *sql.DB

func Init() {
	configs.LoadDBConfig()
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		configs.DBConfig.User,
		configs.DBConfig.User,
		configs.DBConfig.Password,
		configs.DBConfig.Host,
		configs.DBConfig.Port,
		configs.DBConfig.Database,
	)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Could not connect to the databse", err)
	}

	log.Println("Connected to Database!")

	migrationDir := "../migrations"

	if err := goose.Up(DB, migrationDir); err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	log.Println("Migrations ran successfully")

	createUserTable()
	createPostTable()
	createPhotoTable()
}

func createUserTable() {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY, 
		name TEXT NOT NULL, 
		email TEXT UNIQUE NOT NULL, 
		password TEXT NOT NULL, 
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
		);
		`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Created table users")
}

func createPostTable() {
	query := `
	CREATE TABLE IF NOT EXISTS posts (
		id SERIAL PRIMARY KEY, 
		title TEXT NOT NULL, 
		content TEXT NOT NULL, 
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT NULL
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Created table posts")
}

func createPhotoTable() {
	query := `
	CREATE TABLE IF NOT EXISTS photos (
		id SERIAL PRIMARY KEY,
		title TEXT NOT NULL, 
		description TEXT NOT NULL, 
		url TEXT NOT NULL, 
		created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMPTZ DEFAULT NULL
		);
		`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	log.Println("Created table photos")
}
