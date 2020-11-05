package services

import (
	"errors"
	"strconv"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

type accountCreator interface {
	CreateAccount(account *models.Account) (*models.Account, error)
}

type accountService struct{}

// CreateAccount is
func (service accountService) CreateAccount(account *models.Account) (*models.Account, error) {
	if account.DocumentNumber == "" {
		return account, errors.New("DocumentNumber can't be empty")
	}

	if _, err := strconv.Atoi(account.DocumentNumber); err != nil {
		return account, errors.New("DocumentNumber must be a number")
	}

	return repositories.AccountRepository.SaveAccountInDB(account)
}

// AccountService is an instance of the service to be accessed outside
var (
	AccountService accountCreator = accountService{}
)
