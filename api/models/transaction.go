package models

import (
	"time"
)

// Transaction is the struct that defines the transactions.
type Transaction struct {
	// Default columns
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`

	// Custom columns
	AmountCents int       `json:"amount_cents"`
	EventDate   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"eventDate"`

	// Associations
	AccountID       int
	Account         Account
	OperationTypeID int
	OperationType   OperationType
}

// Amount function returns the AmountCents in a formatted value
func (transaction *Transaction) Amount() float64 {
	return float64(transaction.AmountCents) / 100.0
}
