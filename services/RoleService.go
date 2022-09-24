package services

import (
	"simple-crud-golang/models"
)

func FindRoleByName(name string) (*models.Role, bool) {

	var role models.Role
	dbres := db.Model(&models.Role{}).Where("role_name=?", name).First(&role)

	if dbres.RowsAffected == 0 {
		return nil, false
	}
	return &role, true
}

type role struct {
	id   uint
	name string
}

func FindRolesByUser(user User) []models.Role {
	var roles []models.Role
	db.Raw("SELECT roles.id, roles.role_name FROM user_roles right join roles on user_roles.role_id = roles.id where user_roles.user_id = ?", user.ID).Scan(&roles)

	return roles
}
