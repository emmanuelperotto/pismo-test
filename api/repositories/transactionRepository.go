package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
)

// CreateTransaction tries to persist a Transaction in the DB. It can return an error
func CreateTransaction(transaction *models.Transaction) (*models.Transaction, error) {
	// TODO: return better errors
	if err := DB.Create(transaction).Error; err != nil {
		return &models.Transaction{}, err
	}

	return transaction, nil
}
