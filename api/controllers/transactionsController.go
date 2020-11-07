package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/services"
	"github.com/emmanuelperotto/pismo-test/api/utils"
)

// CreateTransaction receives body params and handles transaction creation request
func CreateTransaction(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data models.Transaction
	decoder.Decode(&data)

	transaction, err := services.CreateTransaction.Create(&data)

	if err != nil {
		utils.ErrorResponse(response, http.StatusUnprocessableEntity, err.Error())
	} else {
		utils.JSONResponse(response, http.StatusCreated, transaction)
	}
}
