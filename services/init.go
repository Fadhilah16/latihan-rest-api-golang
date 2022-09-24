package services

import (
	"simple-crud-golang/config"
	"simple-crud-golang/models"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB = config.GetDB()

type Product models.Product
type User models.User
