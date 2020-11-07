package services

import (
	"errors"
	"strconv"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

type accountCreator interface {
	Create(account *models.Account) (*models.Account, error)
}

type createAccountService struct{}

func (service createAccountService) Create(account *models.Account) (*models.Account, error) {
	err := validateAccountFields(account)

	if err != nil {
		return account, err
	}

	return repositories.AccountRepository.SaveAccountInDB(account)
}

// FIXME: add a validation layer and implement validations with Composite Design Pattern
func validateAccountFields(account *models.Account) error {
	// TODO: add validation to return "Account already exists"
	if account.DocumentNumber == "" {
		return errors.New("DocumentNumber can't be empty")
	}

	if _, err := strconv.Atoi(account.DocumentNumber); err != nil {
		return errors.New("DocumentNumber must contain only numbers")
	}

	return nil
}

// CreateAccount is an instance of the service to be accessed outside
var (
	CreateAccount accountCreator = createAccountService{}
)
