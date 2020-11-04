package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/models"
)

// CreateAccount tries to persist an Account in the DB. It can return an error
func CreateAccount(account *models.Account) (*models.Account, error) {
	// TODO: Validate if the documentNumber is a number
	// TODO: return better errors
	if err := DB.Create(account).Error; err != nil {
		return &models.Account{}, err
	}

	return account, nil
}

// GetAccountByID tries to find an account given an ID. It can return an error if not found
func GetAccountByID(id int) (*models.Account, error) {
	var account models.Account

	if err := DB.First(&account, id).Error; err != nil {
		return &account, err
	}

	return &account, nil
}
