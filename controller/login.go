package controller

import (
	"encoding/json"
	"github.com/extark/go_jwt_auth"
	"github.com/extark/jwt_microservice/models"
	"github.com/extark/jwt_microservice/utils"
	"net/http"
	"time"
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
		return
	}

	//Try to validate the obtained values
	if err := utils.Cfg.VALIDATOR.Struct(loginInput); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		return
	}

	//If the credentials are correct get the user from the database
	account, err := models.CheckLogin(utils.Cfg.USERID, loginInput.User, loginInput.Password, utils.Cfg.SQLDB)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		return
	}

	var refresh models.RefreshToken
	// generate a refresh and an access token using go_jwt_auth repo
	refresh.AccessToken, refresh.RefreshToken, err = go_jwt_auth.CreateTokens(account.UUID, utils.Cfg.TOKENEXPIRETIME, utils.Cfg.SECRET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		return
	}

	refresh.ExpireAt = time.Now().Add(time.Hour * 24 * 30)

	// creates a refresh token instance inside the database
	if err = refresh.CreateRefresh(utils.Cfg.SQLDB); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.StandardError{Error: "impossible to create refresh token instance inside the database"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.TokenResponse{UUID: account.UUID, AccessToken: refresh.AccessToken, RefreshToken: refresh.RefreshToken})
	return
}
