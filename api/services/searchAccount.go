package services

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
)

type accountSearcher interface {
	FindByID(id int) (*models.Account, error)
}

type searchAccountService struct{}

func (service searchAccountService) FindByID(id int) (*models.Account, error) {
	return repositories.AccountRepository.GetAccountByID(id)
}

// SearchAccount is an instance of the service to be accessed outside
var (
	SearchAccount accountSearcher = searchAccountService{}
)
