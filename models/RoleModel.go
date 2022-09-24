package models

import "simple-crud-golang/config"

type Role struct {
	Id       int    `json:"id" gorm:"primary_key"`
	RoleName string `json:"role" gorm:"unique"`
	// Users    []User `json:"users" gorm:"many2many:user_roles;"`
}

func init() {
	config.Connect()

	config.GetDB().AutoMigrate(&Role{})
	var countRole int
	config.GetDB().Model(&Role{}).Where("role_name = ?", USER).Count(&countRole)
	if countRole == 0 {
		RoleUser := Role{
			RoleName: "USER",
		}
		config.GetDB().Create(&RoleUser)
	}
	config.GetDB().Model(&Role{}).Where("role_name = ?", ADMIN).Count(&countRole)
	if countRole == 0 {

		RoleAdmin := Role{
			RoleName: "ADMIN",
		}
		config.GetDB().Create(&RoleAdmin)
	}
}
