package services

import (
	"github.com/jinzhu/gorm"
)

func CreateProduct(product *Product) (*Product, *gorm.DB) {
	db.NewRecord(*product)
	dbres := db.Create(product)
	if dbres.Error == nil {

		return product, dbres
	}
	return nil, dbres
}

func UpdateProduct(product *Product) (*Product, *gorm.DB) {
	dbres := db.Save(product)
	if dbres.Error == nil {
		return product, dbres
	}

	return nil, dbres

}

func GetAllProducts() ([]Product, *gorm.DB) {
	var products []Product
	dbres := db.Find(&products)
	return products, dbres
}

func GetProductById(id int64) (*Product, *gorm.DB) {
	var getProduct Product
	db := db.Where("id = ?", id).Find(&getProduct)
	return &getProduct, db
}

func DeleteProduct(id int64) (Product, *gorm.DB) {
	var product Product
	dbres := db.Where("id = ?", id).Delete(product)
	return product, dbres
}
