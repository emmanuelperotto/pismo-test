package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONResponse is a json formatter to be used on controller
func JSONResponse(response http.ResponseWriter, statusCode int, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	err := json.NewEncoder(response).Encode(data)
	if err != nil {
		fmt.Fprintf(response, "%s", err.Error())
	}
}

// ErrorResponse is a json formatter to be used on controllers to standardize error messages
func ErrorResponse(response http.ResponseWriter, statusCode int, message string) {
	JSONResponse(response, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: message,
	})
}
