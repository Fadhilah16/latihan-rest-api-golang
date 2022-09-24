package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	dto "simple-crud-golang/DTO"
	"simple-crud-golang/models"
	"simple-crud-golang/services"
	"simple-crud-golang/utils"

	"github.com/gorilla/mux"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {

	product := &services.Product{}
	utils.ParseBody(r, product)
	p, db := services.CreateProduct(product)
	var response dto.Response
	if db.Error == nil {
		response.Status = http.StatusOK
		response.Message = append(response.Message, "Product successfully created")
		response.Entity = p
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to create product")
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {

	products, db := services.GetAllProducts()
	var status int
	if db.Error == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}

	utils.EncodeJson(w, products, status)

}

func GetProductById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productId := vars["id"]
	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	productDetails, db := services.GetProductById(id)
	var status int
	if db.Error == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}

	utils.EncodeJson(w, productDetails, status)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	var product = &models.Product{}
	utils.ParseBody(r, product)

	var response dto.Response

	if product.Id >= 0 {

		productDetails, _ := services.GetProductById(product.Id)

		if product.Name != "" {
			productDetails.Name = product.Name
		}
		if product.Description != "" {
			productDetails.Description = product.Description
		}
		if product.Price > 0 {
			productDetails.Price = product.Price
		}

		_, db := services.UpdateProduct(productDetails)

		if db.Error == nil {
			response.Status = http.StatusOK
			response.Message = append(response.Message, "Product successfully updated")
			response.Entity = productDetails
		} else {
			response.Status = http.StatusBadRequest
			response.Message = append(response.Message, "Failed to update product")
			response.Message = append(response.Message, db.Error.Error())
			response.Entity = nil
		}

		utils.EncodeJson(w, response, response.Status)
	}

}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productId := vars["id"]
	id, err := strconv.ParseInt(productId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	_, db := services.DeleteProduct(id)
	var response dto.Response
	if db.Error == nil {
		response.Status = http.StatusOK
		response.Message = append(response.Message, "Product successfully deleted")
		response.Entity = nil
	} else {
		response.Status = http.StatusBadRequest
		response.Message = append(response.Message, "Failed to delete product")
		response.Entity = nil
	}

	utils.EncodeJson(w, response, response.Status)
}
