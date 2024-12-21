package main

import (
	"database/sql"
	"log"
	"os"
)

func main() {
	// 1. Get card number input
	// 2. Sanitize input
	// 3. Detect card type
	// 4. Validate using Luhn
	// 5. Store result in database

	// Initialize Db connection
	db, err := sql.Open("mysql", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}
