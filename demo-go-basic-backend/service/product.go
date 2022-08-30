package service

import (
	"demo-go-basic-backend/model"
)

var Products []model.Product

func GetProducts() []model.Product {
	return Products
}

func GetProduct(index int) model.Product {
	return Products[index]
}

func CreateProduct(newProduct model.Product) []model.Product {
	Products = append(Products, newProduct)
	return Products
}

func DeleteProduct(index int) []model.Product {
	Products = append(Products[:index], Products[index + 1:]...)
	return Products
}

func UpdateProduct(index int, newProduct model.Product) []model.Product {
	Products[index] = newProduct
	return Products
}