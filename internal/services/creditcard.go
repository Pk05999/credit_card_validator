package services

import (
	"github.com/Pk05999/credit_card_validator/internal/models"
	"github.com/Pk05999/credit_card_validator/internal/repositories"
)

func ValidateAndCreateCard(userID uint, cardNumber string) (bool, error) {
	isValid := LuhnAlgorithm(cardNumber)

	card := models.CreditCard{
		UserID:     userID,
		CardNumber: cardNumber, // In production, encrypt this!
		IsValid:    isValid,
	}

	if err := repositories.CreateCreditCard(&card); err != nil {
		return false, err
	}

	return isValid, nil
}

// Luhn Algorithm for card validation
func LuhnAlgorithm(card string) bool {
	sum := 0
	double := false
	for i := len(card) - 1; i >= 0; i-- {
		n := int(card[i] - '0')
		if double {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		double = !double
	}
	return sum%10 == 0
}
