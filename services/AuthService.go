package services

import (
	dto "simple-crud-golang/DTO"

	"github.com/jinzhu/gorm"
)

func CreateUser(user *User) (*User, *gorm.DB) {
	_, exists := FindUserByUsername(user.Username)

	if exists == true {

		return nil, nil
	}
	db.NewRecord(*user)
	dbres := db.Create(user)
	return user, dbres

}

func UpdateUser(user *User) (*User, *gorm.DB) {
	_, exists := FindUserByUsername(user.Username)
	var dbres *gorm.DB
	if exists == true {
		dbres = db.Save(user)

		if dbres.Error == nil {
			return user, dbres
		}
	}

	return nil, dbres
}

func FindUserByUsername(username string) (*User, bool) {
	var user User
	dbres := db.Model(&User{}).Where("username=?", username).First(&user)

	if dbres.RowsAffected > 0 {
		return &user, true
	}
	return nil, false
}

func MatchUserProperties(userData dto.Register) *User {
	var user User
	user.Name = userData.Name
	user.Username = userData.Username
	user.Password = userData.Password

	return &user
}
