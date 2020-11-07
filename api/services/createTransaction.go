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
	err := validateTransactionFields(transaction)

	if err != nil {
		return transaction, err
	}

	return repositories.TransactionRepository.SaveTransactionInDB(transaction)
}

// FIXME: add a validation layer and implement validations with Composite Design Pattern
func validateTransactionFields(transaction *models.Transaction) error {
	if transaction.AccountID == 0 {
		return errors.New("AccountID is required")
	}

	if transaction.OperationTypeID == 0 {
		return errors.New("OperationTypeID is required")
	}

	operationType, err := repositories.OperationTypeRepository.GetByID(transaction.OperationTypeID)

	if err != nil {
		return errors.New("OperationTypeID not found")
	}

	if !operationType.ShouldWithdrawBalance && transaction.AmountCents < 0 {
		return fmt.Errorf("AmountCents must be positive when registering: %s", operationType.Description)
	}

	if operationType.ShouldWithdrawBalance && transaction.AmountCents > 0 {
		return fmt.Errorf("AmountCents must be negative when registering: %s", operationType.Description)
	}

	return nil
}

// CreateTransaction is a service that deals with the transaction creation process
var (
	CreateTransaction transactionCreator = createTransaction{}
)
