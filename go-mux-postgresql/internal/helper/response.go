package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Error bool `json:"error"`
	ErrorMessage string `json:"errorMessage"`
}

func SuccessResponse(message string, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Code: 200, Message: message, Data: data, Error: false, ErrorMessage: ""})
}

func FailedResponse(code int, message string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Code: code, Message: "", Data: "null", Error: true, ErrorMessage: message})
}

func ErrorNotFound(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Code: http.StatusNotFound, Message: "", Data: "null", Error: true, ErrorMessage: message})	
}

func MethodNotAllowed(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{Code: http.StatusMethodNotAllowed, Message: "", Data: "null", Error: true, ErrorMessage: message})
}