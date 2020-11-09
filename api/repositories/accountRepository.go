package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
)

type accountRepoInterface interface {
	SaveAccountInDB(account *models.Account) (*models.Account, error)
	GetAccountByID(id int) (*models.Account, error)
	Update(account *models.Account) (*models.Account, error)
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

	err := Repository.DB.First(&account, id).Error

	return &account, err
}

func (repository accountRepo) Update(account *models.Account) (*models.Account, error) {
	err := Repository.DB.Save(account).Error

	return account, err
}

// AccountRepository wraps db queries associated to accounts table
var (
	AccountRepository accountRepoInterface = accountRepo{}
)
