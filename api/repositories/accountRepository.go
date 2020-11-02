package repositories

import (
	"github.com/emmanuelperotto/pismo-test/api/config"
	"github.com/emmanuelperotto/pismo-test/api/models"
)

// TODO: Validate if the documentNumber is a number
func CreateAccount(account *models.Account) (*models.Account, error) {
	// TODO: return better errors
	if err := config.DB.Create(account).Error; err != nil {
		return &models.Account{}, err
	}

	return account, nil
}
