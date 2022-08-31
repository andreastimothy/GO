package model

import (
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
	Error bool `json:"error"`
	ErrorMessage string `json:"error message"`
}

func ErrNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Set Status Code
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(ErrorResponse{Error: true, ErrorMessage: "404: Not Found!"})
}

func ErrMethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Set Status Code
	w.WriteHeader(http.StatusMethodNotAllowed)
	json.NewEncoder(w).Encode(ErrorResponse{Error: true, ErrorMessage: "405: Method Not Allowed!"})
}