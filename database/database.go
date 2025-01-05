package database

import (
	"fmt"
	"log"

	"github.com/Pk05999/credit_card_validator/config"
	"github.com/Pk05999/credit_card_validator/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	cfg := config.GetConfig()

	// Construct the DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connection successful!")

	// Migrate the database schema
	if err := DB.AutoMigrate(&models.CreditCard{}); err != nil {
		log.Fatalf("Failed to migrate database schema: %v", err)
	}
}