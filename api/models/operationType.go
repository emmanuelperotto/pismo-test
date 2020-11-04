package models

import (
	"time"
)

// OperationType is the struct that defines the type of transactions.
type OperationType struct {
	// Default columns
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`

	// Custom columns
	Description string `gorm:"unique;not null" json:"description"`

	// Associations
	Transactions []Transaction `json:"-"`
}
