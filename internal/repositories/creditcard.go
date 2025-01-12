package repositories

import (
	"github.com/Pk05999/credit_card_validator/internal/database"
	"github.com/Pk05999/credit_card_validator/internal/models"
)

func CreateCreditCard(card *models.CreditCard) error {
	return database.DB.Create(card).Error
}
