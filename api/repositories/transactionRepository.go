package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
)

// TODO: add tests

type transactionRepoInterface interface {
	SaveTransactionInDB(transaction *models.Transaction) (*models.Transaction, error)
}

type transactionRepo struct{}

// SaveTransactionInDB tries to persist a Transaction in the DB. It can return an error
func (repository transactionRepo) SaveTransactionInDB(transaction *models.Transaction) (*models.Transaction, error) {
	err := Repository.DB.Create(transaction).Error
	return transaction, err
}

// TransactionRepository is a repository that wraps queries for "transactions" table
var (
	TransactionRepository transactionRepoInterface = transactionRepo{}
)
