package handlers

import (
	"net/http"

	"github.com/Pk05999/credit_card_validator/internal/services"
	"github.com/gin-gonic/gin"
)

func AddCreditCard(c *gin.Context) {
	var input struct {
		UserID     uint   `json:"user_id" binding:"required"`
		CardNumber string `json:"card_number" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isValid, err := services.ValidateAndCreateCard(input.UserID, input.CardNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Card added successfully", "is_valid": isValid})
}

func GetAllCards(c *gin.Context) {
	cards, err := services.GetAllCards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cards)
}
