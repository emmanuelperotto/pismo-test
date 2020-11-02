package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	"github.com/emmanuelperotto/pismo-test/api/utils"
)

func CreateAccount(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data models.Account
	decoder.Decode(&data)

	acc, err := repositories.CreateAccount(&data)

	if err != nil {
		utils.ErrorResponse(response, http.StatusUnprocessableEntity, errors.New("Invalid document number").Error())
	} else {
		utils.JSONResponse(response, http.StatusCreated, acc)
	}
}
