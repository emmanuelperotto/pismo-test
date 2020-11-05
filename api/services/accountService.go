package services

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

type accountCreator interface {
	CreateAccount(account *models.Account) (*models.Account, error)
}

type accountService struct{}

// CreateAccount is
func (service accountService) CreateAccount(account *models.Account) (*models.Account, error) {
	return repositories.AccountRepository.SaveAccountInDB(account)
}

// AccountService is an instance of the service to be accessed outside
var (
	AccountService accountCreator = accountService{}
)
