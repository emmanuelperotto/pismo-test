package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/emmanuelperotto/pismo-test/api/config"
	"github.com/emmanuelperotto/pismo-test/api/controllers"
	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// TODO: extract to small functions

	// Connect to DB
	var err error
	err = godotenv.Load()
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbConfig := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)
	config.DB, err = gorm.Open(postgres.Open(dbConfig), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	config.DB.AutoMigrate(&models.Account{}, &models.OperationType{})
	config.SeedDB()

	// Define routes
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")

	// Start server
	address := ":3000"
	srv := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Running on " + address)
	log.Fatal(srv.ListenAndServe())
}
