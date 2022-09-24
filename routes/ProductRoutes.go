package routes

import (
	"simple-crud-golang/controllers"
	"simple-crud-golang/middleware"
	"simple-crud-golang/models"

	"github.com/gorilla/mux"
)

var RegisterProductRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/products/", middleware.Middleware(controllers.CreateProduct, models.ADMIN)).Methods("POST")
	router.HandleFunc("/api/products/", middleware.Middleware(controllers.GetProducts, models.USER)).Methods("GET")
	router.HandleFunc("/api/products/{id}", middleware.Middleware(controllers.GetProductById, models.USER)).Methods("GET")
	router.HandleFunc("/api/products/", middleware.Middleware(controllers.UpdateProduct, models.ADMIN)).Methods("PUT")
	router.HandleFunc("/api/products/{id}", middleware.Middleware(controllers.DeleteProduct, models.ADMIN)).Methods("DELETE")

}
