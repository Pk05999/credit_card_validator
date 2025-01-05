package models

import (
	"gorm.io/gorm"
)

type CreditCard struct {
	gorm.Model
	CardNumber string `gorm:"unique;not null" json:"card_number" binding:"required,credit_card"`
	CardHolder string `gorm:"not null" json:"card_holder" binding:"required"`
	ExpiryDate string `gorm:"not null" json:"expiry_date" binding:"required"`
}
