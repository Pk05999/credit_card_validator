package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
}

var config Config

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config = Config{
		Port:   getEnv("PORT", "8080"),
		DbUser: getEnv("DB_USER", "root"),
		DbPass: getEnv("DB_PASSWORD", ""),
		DbHost: getEnv("DB_HOST", "localhost"),
		DbPort: getEnv("DB_PORT", "3306"),
		DbName: getEnv("DB_NAME", "credit_card_db"),
	}
}

func GetConfig() Config {
	return config
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
