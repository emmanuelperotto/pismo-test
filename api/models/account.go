package models

import (
	"time"
)

// Account is the struct that defines the account model
type Account struct {
	ID             uint      `gorm:"primaryKey" json:"id"`
	DocumentNumber string    `gorm:"unique;not null" json:"documentNumber"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
