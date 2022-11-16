package models

type StandardError struct {
	Error string `json:"error"`
}

type TokenResponse struct {
	UUID         string `json:"uuid"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
