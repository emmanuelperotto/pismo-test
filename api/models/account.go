package models

import (
	"time"
)

// Account is the struct that defines the account model
type Account struct {
	// Default columns
	ID        int       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`

	// Custom columns
	DocumentNumber            string `gorm:"unique;not null" json:"documentNumber"`
	AvailableCreditLimitCents int    `gorm:"not null" json:"availableCreditLimitCents"`

	// Associations
	Transactions []Transaction `json:"-"`
}
