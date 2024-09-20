package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	connStr := "postgres://postgres:Andreanrue1!@localhost:5432/matthewrowley?sslmode=disable"
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

	createTable()
}

func createTable() {
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
