package config

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emmanuelperotto/pismo-test/api/controllers"
	"github.com/gorilla/mux"
)

// RunServer setup the routes, config the server and run it
func RunServer() {
	// Define routes
	router := mux.NewRouter()
	router.HandleFunc("/accounts/{id}", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/accounts", controllers.CreateAccount).Methods("POST")

	// Start server
	address := ":3000"
	server := &http.Server{
		Handler:      router,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Running on " + address)
	log.Fatal(server.ListenAndServe())
}
