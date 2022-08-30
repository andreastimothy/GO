package route

import (
	"github.com/gorilla/mux"
	"net/http"
	"demo-go-basic-backend/handler"
)

func InitProductRoute(r *mux.Router) {
	//GET PRODUCT BY ID
	r.HandleFunc("/products/{id}", handler.GetProductHandler).Methods(http.MethodGet)

	//CREATE PRODUCT
	r.HandleFunc("/products", handler.CreateProductHandler).Methods(http.MethodPost)

	//GET ALL PRODUCTS
	r.HandleFunc("/products", handler.GetProductsHandler).Methods(http.MethodGet)

	//DELETE PRODUCT
	r.HandleFunc("/products/{id}", handler.DeleteProductHandler).Methods(http.MethodDelete)

	//UPDATE PRODUCT
	r.HandleFunc("/products/{id}", handler.UpdateProductHandler).Methods(http.MethodPut)
}