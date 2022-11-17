package controller

import (
	"encoding/json"
	"fmt"
	"github.com/extark/go_jwt_auth"
	"github.com/extark/jwt_microservice/models"
	"github.com/extark/jwt_microservice/utils"
	"log"
	"net/http"
	"time"
)

// Refresh when the access token is expired this function refresh with a second token the first token
func Refresh(w http.ResponseWriter, r *http.Request) {
	log.Println("GET - /refresh")

	//Set headers content type
	w.Header().Set("Content-Type", "application/json")

	//Try to decode the body
	var tokens models.TokenResponse
	if err := json.NewDecoder(r.Body).Decode(&tokens); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		log.Panicln(err.Error())
		return
	}

	//Try to validate the obtained values
	if err := utils.Cfg.VALIDATOR.Struct(tokens); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		log.Panicln(err.Error())
		return
	}

	// Check if the passed token is valid and if contains a data
	accountUuid, err := go_jwt_auth.GetTokenData(tokens.RefreshToken, utils.Cfg.SECRET)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		log.Panicln(err.Error())
		return
	}

	// Check if the refresh token exists inside the database
	refresh := models.RefreshToken{
		RefreshToken: tokens.RefreshToken,
		AccessToken:  tokens.AccessToken,
	}

	if ok, err := refresh.IsRefreshTokenValid(utils.Cfg.SQLDB); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		log.Panicln(err.Error())
		return
	} else if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.StandardError{Error: "impossible to find this record inside the database"})
		log.Panicln(err.Error())
		return
	}

	//If the token is valid and is inside the database, it removes the old row
	if err := refresh.DeleteRefreshToken(utils.Cfg.SQLDB); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.StandardError{Error: "impossible to delete old refresh token instance"})
		return
	}

	//Generate new access and refresh token
	accessT, refreshT, err := go_jwt_auth.CreateTokens(fmt.Sprintf("%v", accountUuid), utils.Cfg.TOKENEXPIRETIME, utils.Cfg.SECRET)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(models.StandardError{Error: err.Error()})
		log.Panicln(err.Error())
		return
	}

	refresh.AccessToken = accessT
	refresh.RefreshToken = refreshT
	refresh.ExpireAt = time.Now().Add(time.Hour * 24 * 30)

	//Try to create new instance inside access-refresh token table
	if err = refresh.CreateRefresh(utils.Cfg.SQLDB); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.StandardError{Error: "impossible to create refresh token instance on DB"})
		log.Panicln(err.Error())
		return
	}

	//If all are ok returns the new tokens
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.TokenResponse{UUID: fmt.Sprintf("%v", accountUuid), AccessToken: accessT, RefreshToken: refreshT})
}
