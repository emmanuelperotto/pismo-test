package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/repositories"
	"github.com/emmanuelperotto/pismo-test/api/utils"
	"github.com/gorilla/mux"
)

// CreateAccount receives body params and creates an Account
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

// GetAccount tries to find an Account given an ID through query params
func GetAccount(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	acc, err := repositories.GetAccountByID(id)
	if err != nil {
		utils.ErrorResponse(response, http.StatusNotFound, err.Error())
	} else {
		utils.JSONResponse(response, http.StatusOK, acc)
	}
}
