package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/emmanuelperotto/pismo-test/api/models"
	"github.com/emmanuelperotto/pismo-test/api/services"
	"github.com/emmanuelperotto/pismo-test/api/utils"
	"github.com/gorilla/mux"
)

// CreateAccount receives body params and creates an Account
func CreateAccount(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var data models.Account
	decoder.Decode(&data)

	account, err := services.CreateAccount.Create(&data)

	if err != nil {
		utils.ErrorResponse(response, http.StatusUnprocessableEntity, err.Error())
	} else {
		utils.JSONResponse(response, http.StatusCreated, account)
	}
}

// GetAccount tries to find an Account given an ID through query params
func GetAccount(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := strconv.Atoi(params["id"])

	account, err := services.SearchAccount.FindByID(id)
	if err != nil {
		utils.ErrorResponse(response, http.StatusNotFound, err.Error())
	} else {
		utils.JSONResponse(response, http.StatusOK, account)
	}
}
