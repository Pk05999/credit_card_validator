package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv" // To load .env file
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve environment variables
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN is not set in the environment")
	}

	// Initialize DB connection
	db, err := sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Test the DB connection
	if err = db.Ping(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Println("Database connection successful!")
}
