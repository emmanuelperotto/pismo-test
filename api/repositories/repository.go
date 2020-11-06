package repositories

import (
	"fmt"
	"os"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

var (
	Repository Repo
)

// SetupDB is the function that initializes the DB connecting to it, migrating and seeding data
func SetupDB() {
	err := godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	Repository.DB, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	Repository.DB.AutoMigrate(&models.Account{}, &models.OperationType{}, &models.Transaction{})
	seedDB()
}

func seedDB() {
	operationTypes := []models.OperationType{
		{
			Description:           "COMPRA À VISTA",
			ShouldWithdrawBalance: true,
		},
		{
			Description:           "COMPRA PARCELADA",
			ShouldWithdrawBalance: true,
		},
		{
			Description:           "SAQUE",
			ShouldWithdrawBalance: true,
		},
		{
			Description:           "PAGAMENTO",
			ShouldWithdrawBalance: false,
		},
	}

	for _, operationType := range operationTypes {
		fmt.Println("Trying to find or create OperationType with description: " + operationType.Description)
		err := Repository.DB.Where("description = ?", operationType.Description).FirstOrCreate(&operationType).Error

		if err != nil {
			fmt.Println("Error: " + err.Error())
		}
	}
}
