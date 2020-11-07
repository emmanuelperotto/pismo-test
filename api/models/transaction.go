package models

import (
	"time"
)

// Transaction is the struct that defines the transactions.
type Transaction struct {
	// Default columns
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"-"`

	// Custom columns
	AmountCents int       `json:"amountCents"`
	EventDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"eventDate"`

	// Associations
	AccountID       int           `json:"accountID"`
	Account         Account       `json:"-"`
	OperationTypeID int           `json:"operationTypeID"`
	OperationType   OperationType `json:"-"`
}
