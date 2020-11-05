package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
)

type accountRepoInterface interface {
	SaveAccountInDB(account *models.Account) (*models.Account, error)
}

type accountRepo struct{}

// SaveAccountInDB tries to persist an Account in the DB. It can return an error
func (repository accountRepo) SaveAccountInDB(account *models.Account) (*models.Account, error) {
	err := Repository.DB.Create(account).Error

	return account, err
}

// GetAccountByID tries to find an account given an ID. It can return an error if not found
func (repository accountRepo) GetAccountByID(id int) (*models.Account, error) {
	var account models.Account

	if err := Repository.DB.First(&account, id).Error; err != nil {
		return &account, err
	}

	return &account, nil
}

// AccountRepository is
var (
	AccountRepository accountRepoInterface = accountRepo{}
)
