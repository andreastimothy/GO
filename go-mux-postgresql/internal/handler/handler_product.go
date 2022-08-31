package handler

import (
	"database/sql"
	"demo-go-basic-backend/internal/model"
	"demo-go-basic-backend/internal/helper"
	"demo-go-basic-backend/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetProductsHandler(w http.ResponseWriter, r *http.Request) {
	qSearch := r.URL.Query().Get("search")
	qSort := r.URL.Query().Get("sort")
	qPage := r.URL.Query().Get("page")
	qSize := r.URL.Query().Get("size")

	qSearch = "%" + qSearch + "%"

	var page, size int
	var err error

	if qSort == "" && qSort != "DESC" {
		qSort = "ASC"
	}
	if qPage == "" {
		page = 1
	}
	if qSize == "" {
		size = 10
	}
	if qPage != "" {
		page, err = strconv.Atoi(qPage)
		if err != nil {
			helper.FailedResponse(http.StatusBadRequest, err.Error(), w)
			return
		}
	}
	if qSize != "" {
		size, err = strconv.Atoi(qSize)
		if err != nil {
			helper.FailedResponse(http.StatusBadRequest, err.Error(), w)
		}
	}

	result, err := service.GetProducts(page, size, qSearch, qSort)
	if err != nil {
		helper.FailedResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}
	helper.SuccessResponse(http.StatusText(http.StatusOK), result, w)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		helper.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		return
	}

	result, err := service.GetProduct(id)
	if err == sql.ErrNoRows {
		helper.FailedResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
		return
	}

	if err != nil {
		helper.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		return
	}
	helper.SuccessResponse(http.StatusText(http.StatusOK), result, w)
}

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product *model.Product

	//decode json from request to golang
	json.NewDecoder(r.Body).Decode(&product)

	//validation
	if product.Name == "" {
		helper.FailedResponse(http.StatusBadRequest, "Missing Name Field", w)
		return
	}

	if product.Description == "" {
		helper.FailedResponse(http.StatusBadRequest, "Missing Description Field", w)
		return
	}

	if product.Quantity == 0 {
		helper.FailedResponse(http.StatusBadRequest, "Missing Quantity Field", w)
		return
	}

	result, err := service.CreateProduct(product.Name, product.Description, product.Quantity)
	if err != nil {
		helper.FailedResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	helper.SuccessResponse("Successfully Create Product", result, w)
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		helper.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		return
	}

	_, err = service.DeleteProduct(id)
	if err == sql.ErrNoRows {
		helper.FailedResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
		return
	}
	if err != nil {
		helper.FailedResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
		return
	}
	helper.SuccessResponse("Successfully Delete Product", id, w)
}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		helper.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		return
	}

	body := &model.Product{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(body)
	if err != nil {
		helper.FailedResponse(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), w)
		return
	}

	//validation
	if body.Name == "" {
		helper.FailedResponse(http.StatusBadRequest, "Missing Name Field", w)
		return
	}

	if body.Description == "" {
		helper.FailedResponse(http.StatusBadRequest, "Missing Description Field", w)
		return
	}

	if body.Quantity == 0 {
		helper.FailedResponse(http.StatusBadRequest, "Missing Quantity Field", w)
		return
	}

	_, err = service.UpdateProduct(id, body)
	if err == sql.ErrNoRows {
		helper.FailedResponse(http.StatusNotFound, http.StatusText(http.StatusNotFound), w)
		return
	}

	if err != nil {
		helper.FailedResponse(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), w)
		return
	}

	helper.SuccessResponse("Successfully Update Product", nil, w)
}