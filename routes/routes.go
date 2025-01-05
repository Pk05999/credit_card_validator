package routes

import (
	"github.com/Pk05999/credit_card_validator/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/cards", controllers.AddCard)
		api.GET("/cards", controllers.GetAllCards)
	}
}
