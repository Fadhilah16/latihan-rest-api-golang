package controllers

import (
	"net/http"
	dto "simple-crud-golang/DTO"
	"simple-crud-golang/models"
	"simple-crud-golang/services"
	"simple-crud-golang/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	userData := &dto.Register{}
	utils.ParseBody(r, userData)

	user := &services.User{}
	user = services.MatchUserProperties(*userData)

	role, _ := services.FindRoleByName(models.USER)

	user.Roles = append(user.Roles, *role)

	user.Password, _ = utils.GenerateHashPassword(user.Password)
	u, db := services.CreateUser(user)
	var response dto.Response
	if u == nil {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Username has been taken")
		response.Entity = nil
	} else {
		if db.Error == nil {
			response.Status = http.StatusOK
			response.Message = append(response.Message, "You are already registered!")
			response.Entity = u
		} else {
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Failed to register")
			response.Message = append(response.Message, db.Error.Error())
			response.Entity = nil
		}

	}

	utils.EncodeJson(w, response, response.Status)

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	loginUser := &dto.Authentication{}
	utils.ParseBody(r, loginUser)
	authUser, exists := services.FindUserByUsername(loginUser.Username)
	if !exists {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Username or Password is Incorrect")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	passwordIsMatched := utils.CheckPasswordHash(loginUser.Password, authUser.Password)

	if !passwordIsMatched {
		var response dto.Response
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Username or Password is Incorrect")
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}
	authUser.Roles = services.FindRolesByUser(*authUser)
	token, err := utils.GenerateJWT(authUser.Username, authUser.Roles)
	if err != nil {
		var response dto.Response
		response.Status = http.StatusInternalServerError
		response.Message = append(response.Message, "Something went wrong!")
		response.Message = append(response.Message, err.Error())
		response.Entity = nil
		utils.EncodeJson(w, response, response.Status)
		return
	}

	var tokenData dto.Token
	tokenData.Username = authUser.Username
	tokenData.Roles = authUser.Roles
	tokenData.Token = token
	var response dto.Response
	response.Status = http.StatusOK
	response.Message = append(response.Message, "You are already logged in")
	response.Entity = tokenData
	utils.EncodeJson(w, response, response.Status)

}

func AddAdminRole(w http.ResponseWriter, r *http.Request) {
	authUser, exists := services.FindUserByUsername(r.Header.Get("username"))
	var response dto.Response
	if !exists {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "User doesn't exist")
		response.Entity = nil
	} else {
		adminRole, _ := services.FindRoleByName(models.ADMIN)
		roles := services.FindRolesByUser(*authUser)
		authUser.Roles = roles
		authUser.Roles = append(authUser.Roles, *adminRole)
		_, dbres := services.UpdateUser(authUser)

		if dbres.Error == nil {
			response.Status = http.StatusOK
			response.Message = append(response.Message, "Successfully updated user account")
			response.Entity = authUser
		} else {
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Failed to update user account")
			response.Message = append(response.Message, dbres.Error.Error())
			response.Entity = nil
		}
	}
	utils.EncodeJson(w, response, response.Status)
}

func GetUserData(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// username := vars["username"]
	query := r.URL.Query()
	username := query["username"]

	var user *services.User
	var exists bool

	user, exists = services.FindUserByUsername(username[0])

	var response dto.Response
	if !exists {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "User doesn't exist")
		response.Entity = nil
	} else {
		roles := services.FindRolesByUser(*user)
		user.Roles = roles

		response.Status = http.StatusOK
		response.Message = nil
		response.Entity = user

	}
	utils.EncodeJson(w, response, response.Status)
}

func GetSelfUserData(w http.ResponseWriter, r *http.Request) {

	var user *services.User
	var exists bool

	user, exists = services.FindUserByUsername(r.Header.Get("username"))

	var response dto.Response
	if !exists {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "User doesn't exist")
		response.Entity = nil
	} else {
		roles := services.FindRolesByUser(*user)
		user.Roles = roles

		response.Status = http.StatusOK
		response.Message = nil
		response.Entity = user

	}
	utils.EncodeJson(w, response, response.Status)
}
