package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	"github.com/emmanuelperotto/pismo-test/api/utils"
)

// CreateTransaction receives body params and creates a Transaction
func CreateTransaction(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data models.Transaction
	decoder.Decode(&data)

	transaction, err := repositories.Repository.CreateTransaction(&data)

	if err != nil {
		utils.ErrorResponse(response, http.StatusUnprocessableEntity, err.Error())
	} else {
		utils.JSONResponse(response, http.StatusCreated, transaction)
	}
}
