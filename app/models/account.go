package models

import (
	"time"
)

type Account struct {
	ID             uint      `gorm:"primaryKey"`
	CreatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DocumentNumber string    `gorm:"unique;not null"`
}
