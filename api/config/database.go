package config

import (
	"fmt"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"gorm.io/gorm"
)

// SeedDB is a function that initializes my database with some data.
func SeedDB() {
	operationTypes := []models.OperationType{
		{
			Description: "COMPRA À VISTA",
		},
		{
			Description: "COMPRA PARCELADA",
		},
		{
			Description: "SAQUE",
		},
		{
			Description: "PAGAMENTO",
		},
	}

	for _, operationType := range operationTypes {
		fmt.Println("Trying to find or create OperationType with description: " + operationType.Description)
		err := DB.Where("description = ?", operationType.Description).FirstOrCreate(&operationType).Error

		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}
}

// DB is a global instance to access database
var (
	DB *gorm.DB
)
