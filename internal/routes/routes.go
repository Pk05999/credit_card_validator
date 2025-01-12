package routes

import (
	"github.com/Pk05999/credit_card_validator/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/cards", handlers.AddCreditCard)
		// api.GET("/cards", handlers.GetAllCards)
	}
}
