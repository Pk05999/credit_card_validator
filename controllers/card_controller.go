package controllers

import (
	"net/http"

	"github.com/Pk05999/credit_card_validator/database"
	"github.com/Pk05999/credit_card_validator/models"
	"github.com/gin-gonic/gin"
)

func AddCard(c *gin.Context) {
	var card models.CreditCard
	if err := c.ShouldBindJSON(&card); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&card).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save card"})
		return
	}

	c.JSON(http.StatusCreated, card)
}

func GetAllCards(c *gin.Context) {
	var cards []models.CreditCard
	if err := database.DB.Find(&cards).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cards"})
		return
	}

	c.JSON(http.StatusOK, cards)
}
