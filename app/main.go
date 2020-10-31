package main

import (
	"fmt"

	"github.com/emmanuelperotto/pismo-test/app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Starting...")
	dsn := "user=postgres password=secret123 dbname=pismo_development port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.Account{})
}
