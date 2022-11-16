package models

type LoginInput struct {
	User     string `json:"user" validator:"required"`
	Password string `json:"password" validator:"required"`
}
