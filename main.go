package main

import (
	"log"

	"github.com/Pk05999/credit_card_validator/config"
	"github.com/Pk05999/credit_card_validator/internal/database"
	"github.com/Pk05999/credit_card_validator/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//Load eenvironment variables
	// config.LoadConfig()
	config.GetConfig()

	//Initialize database
	database.ConnectDatabase()

	// Initialize Gin router
	r := gin.Default()

	//Setup routes
	routes.SetupRoutes(r)

	//Start the server
	// port := config.GetConfig().Port
	// log.Printf("Server started on port %s", port)
	// r.Run(":" + port)

	// Start the server
	r.Run(":8080")
	log.Println("Server started......")

}
