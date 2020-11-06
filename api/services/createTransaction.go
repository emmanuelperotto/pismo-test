package services

import (
	"errors"
	"fmt"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

type transactionCreator interface {
	Create(transaction *models.Transaction) (*models.Transaction, error)
}

type createTransaction struct{}

func (service createTransaction) Create(transaction *models.Transaction) (*models.Transaction, error) {
	if transaction.AccountID == 0 {
		return transaction, errors.New("AccountID is required")
	}

	if transaction.OperationTypeID == 0 {
		return transaction, errors.New("OperationTypeID is required")
	}

	operationType, err := repositories.OperationTypeRepository.GetByID(transaction.OperationTypeID)

	if err != nil {
		return transaction, errors.New("OperationTypeID not found")
	}

	if !operationType.ShouldWithdrawBalance && transaction.AmountCents < 0 {
		return transaction, fmt.Errorf("AmountCents must be positive when registering: %s", operationType.Description)
	}

	if operationType.ShouldWithdrawBalance && transaction.AmountCents > 0 {
		return transaction, fmt.Errorf("AmountCents must be negative when registering: %s", operationType.Description)
	}

	return repositories.TransactionRepository.SaveTransactionInDB(transaction)
}

// CreateTransaction is a service that deals with the transaction creation process
var (
	CreateTransaction transactionCreator = createTransaction{}
)
