package routes

import (
	"simple-crud-golang/controllers"
	"simple-crud-golang/middleware"
	"simple-crud-golang/models"

	"github.com/gorilla/mux"
)

var RegisterAuthRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/auth/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/api/auth/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/api/auth/add-admin-role", middleware.Middleware(controllers.AddAdminRole, models.USER)).Methods("POST")
	router.HandleFunc("/api/auth/account/me", middleware.Middleware(controllers.GetSelfUserData, models.USER)).Methods("GET")
	router.HandleFunc("/api/auth/account", controllers.GetUserData).Methods("GET")
}
