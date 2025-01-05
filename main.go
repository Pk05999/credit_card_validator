package main

import (
	"log"

	"github.com/Pk05999/credit_card_validator/config"
	"github.com/Pk05999/credit_card_validator/database"
	"github.com/Pk05999/credit_card_validator/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	//Load eenvironment variables
	// config.LoadConfig()
	config.LoadConfig()

	//Initialize database
	database.ConnectDatabase()

	//Create Gin router
	router := gin.Default()

	//Setup routes
	routes.RegisterRoutes(router)

	//Start the server
	port := config.GetConfig().Port
	log.Printf("Server started on port %s", port)
	router.Run(":" + port)

}
