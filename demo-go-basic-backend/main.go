package main

import (
	"demo-go-basic-backend/helper"
	"demo-go-basic-backend/model"
	"demo-go-basic-backend/service"
	"demo-go-basic-backend/route"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)
	
func main() {

	r := mux.NewRouter()

	//SEEDING DATA/APPEND DATA
	service.Products = append(service.Products, model.Product{ID: 1, Name: "Pampers Hokage", Description: "for daily use", Quantity: 5})
	service.Products = append(service.Products, model.Product{ID: 2, Name: "Sarung Tsunade", Description: "Not for Child", Quantity: 10})
	service.Products = append(service.Products, model.Product{ID: 3, Name: "Landyard Uchiha", Description: "Not for Itachi", Quantity: 7})

	// productRouter := r.PathPrefix("model.Products").Subrouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "API Warung 1.0")
	}).Methods(http.MethodGet)

	r.Use(helper.LogMiddleware)

	r.NotFoundHandler = http.HandlerFunc(model.ErrNotFound)
	r.MethodNotAllowedHandler = http.HandlerFunc(model.ErrMethodNotAllowed)

	route.InitRoute(r)

	fmt.Println("Starting web server on 8080")
	http.ListenAndServe(":8080", r)
}