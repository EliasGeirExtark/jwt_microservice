package controller

import (
	"encoding/json"
	"github.com/extark/jwt_microservice/models"
	"github.com/extark/jwt_microservice/utils"
	"net/http"
)

// Login this function gets a username and a password, checks this values and if are correct create a token and refresh token
func Login(w http.ResponseWriter, r *http.Request) {
	//Set headers content type
	w.Header().Set("Content-Type", "application/json")

	//Try to decode the body
	var loginInput models.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
	}

	//Try to validate the obtained values
	if err := utils.Cfg.VALIDATOR.Struct(loginInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
	}

}
