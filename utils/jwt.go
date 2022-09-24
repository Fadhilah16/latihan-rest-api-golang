package utils

import (
	"fmt"
	"simple-crud-golang/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const SECRET_KEY = "janganlupabobok"

func GenerateJWT(username string, roles []models.Role) (string, error) {

	var mySigningKey = []byte(SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	var claimRoles []string

	for _, role := range roles {
		claimRoles = append(claimRoles, role.RoleName)
	}

	claims["authorized"] = true
	claims["username"] = username
	claims["roles"] = claimRoles
	claims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Printf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
