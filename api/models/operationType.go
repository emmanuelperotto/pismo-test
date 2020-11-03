package models

import (
	"time"
)

// OperationType is the struct that defines the type of transactions.
type OperationType struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Description string    `gorm:"unique;not null" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
