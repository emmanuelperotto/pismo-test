package services

import (
	"errors"
	"strconv"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

// FIXME: Extract it to 2 services/usecases (Creator and Reader)
type accountCreator interface {
	CreateAccount(account *models.Account) (*models.Account, error)
	FindAccountByID(id int) (*models.Account, error)
}

type accountService struct{}

func (service accountService) CreateAccount(account *models.Account) (*models.Account, error) {
	if account.DocumentNumber == "" {
		return account, errors.New("DocumentNumber can't be empty")
	}

	if _, err := strconv.Atoi(account.DocumentNumber); err != nil {
		return account, errors.New("DocumentNumber must be a number")
	}

	// TODO: add validation to return "Account already exists"

	return repositories.AccountRepository.SaveAccountInDB(account)
}

func (service accountService) FindAccountByID(id int) (*models.Account, error) {
	return repositories.AccountRepository.GetAccountByID(id)
}

// AccountService is an instance of the service to be accessed outside
var (
	AccountService accountCreator = accountService{}
)
