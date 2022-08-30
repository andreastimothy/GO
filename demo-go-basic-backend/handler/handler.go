package handler

import (
	"net/http"
	"encoding/json"
	"demo-go-basic-backend/service"
	"demo-go-basic-backend/model"
	"github.com/gorilla/mux"
	"strconv"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	model.SuccessResponse("Successfully Get All Products", service.GetProducts(), w)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		model.FailedResponseBadRequest("Bad Request", w)
		return
	}

	for index, product := range service.Products {
		if product.ID == id {
			model.SuccessResponse("Successfully Get Product By ID", service.GetProduct(index), w)
			return
		}
	}
	model.ErrorNotFound("Product Not Found", w)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product model.Product

	//decode json from request to golang
	json.NewDecoder(r.Body).Decode(&product)

	//validation
	if product.Name == "" {
		model.FailedResponseBadRequest("Missing Name Field", w)
		return
	}

	if product.Description == "" {
		model.FailedResponseBadRequest("Missing Description Field", w)
		return
	}

	if product.Quantity == 0 {
		model.FailedResponseBadRequest("Missing Quantity Field", w)
		return
	}

	//generate ID
	product.ID = len(service.Products) + 1

	//push to product slice
	model.SuccessResponse("Successfully Create Product", service.CreateProduct(product), w)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		model.FailedResponseBadRequest("Bad Request", w)
		return
	}

	for index, product := range service.Products {
		if product.ID == id {
			model.SuccessResponse("Successfully Delete Product", service.DeleteProduct(index), w)
			return
		}
	}
	model.ErrorNotFound("Product Not Found", w)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		model.FailedResponseBadRequest("Bad Request", w)
		return
	}

	for index, product := range service.Products {
		if product.ID == id {
			var product model.Product

			//decode json from request to golang
			json.NewDecoder(r.Body).Decode(&product)

			//validation
			if product.Name == "" {
				model.FailedResponseBadRequest("Missing Name Field", w)
				return
			}

			if product.Description == "" {
				model.FailedResponseBadRequest("Missing Description Field", w)
				return
			}

			if product.Quantity == 0 {
				model.FailedResponseBadRequest("Missing Quantity Field", w)
				return
			}

			//generate ID
			product.ID = id

			//push to product slice
			model.SuccessResponse("Successfully Update Product", service.UpdateProduct(index, product), w)
			return
		}
	}
	model.FailedResponseBadRequest("Bad Request", w)
}