package models

import (
	"simple-crud-golang/config"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"-"`
	Roles    []Role `json:"roles" gorm:"many2many:user_roles;"`
}

func init() {
	config.Connect()
	config.GetDB().AutoMigrate(&User{})
}
