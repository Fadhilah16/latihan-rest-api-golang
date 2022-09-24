package middleware

import (
	"fmt"
	"net/http"
	dto "simple-crud-golang/DTO"
	"simple-crud-golang/services"
	"simple-crud-golang/utils"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func Middleware(h http.HandlerFunc, role string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			var response dto.Response
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Token not found")
			response.Entity = nil
			utils.EncodeJson(w, response, response.Status)
			return
		}

		authorizationHeader := r.Header.Get("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			var response dto.Response
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Invalid token")
			response.Entity = nil
			utils.EncodeJson(w, response, response.Status)
			return
		}

		if authorizationHeader == "" {
			var response dto.Response
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Your token is empty")
			response.Entity = nil
			utils.EncodeJson(w, response, response.Status)
			return
		}

		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		var secretkey = []byte(utils.SECRET_KEY)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return secretkey, nil
		})

		if err != nil {
			var response dto.Response
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Your Token has been expired")
			response.Message = append(response.Message, err.Error())
			response.Entity = nil
			utils.EncodeJson(w, response, response.Status)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user, _ := services.FindUserByUsername(claims["username"].(string))

			roles := services.FindRolesByUser(*user)

			var rolesString []string
			for _, v := range roles {
				rolesString = append(rolesString, v.RoleName)
			}

			if utils.ValueInsideOfSlice(role, rolesString) {

				r.Header.Set("username", claims["username"].(string))
				h.ServeHTTP(w, r)
				return
			}

		}
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Not Authorized")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
	})

}

func AuthChecker(role string, allowedRole string) bool {
	var w http.ResponseWriter
	if role != allowedRole {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Not Authorized")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return false
	}
	return true
}
