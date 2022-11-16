package models

type LoginInputEmail struct {
	User     string `json:"user" validator:"required"`
	Password string `json:"password" validator:"required"`
}
