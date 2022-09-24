package models

import (
	"simple-crud-golang/config"
)

type Product struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func init() {
	config.Connect()
	config.GetDB().AutoMigrate(&Product{})
}
