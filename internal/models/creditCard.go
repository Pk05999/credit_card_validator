package models

type CreditCard struct {
    ID         uint   `gorm:"primaryKey"`
    UserID     uint   `gorm:"not null"`
    CardNumber string `gorm:"size:19;not null"` // Store encrypted in production
    IsValid    bool   `gorm:"not null"`
}
