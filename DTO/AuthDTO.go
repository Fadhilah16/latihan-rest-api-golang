package dto

import "simple-crud-golang/models"

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Username string        `json:"username"`
	Roles    []models.Role `json:"roles"`
	Token    string        `json:"token"`
}

type Register struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
