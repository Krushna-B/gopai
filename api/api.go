package api

import (
	"encoding/json"
	"net/http"
)

// Coint Balance Paras
type CoinBalanceParams struct {
	Username string
}

// Response from server
type CoinBalanceResponse struct {
	//Sucess code
	Code    int
	Balance uint64
}

// Error Reponse
type Error struct {
	//Error code
	Code int
	//Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder().Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHanlder = func(w http.ResponseWriter) {
		writeError(w, "Unexpected Error Occurrd", http.StatusInternalServerError)
	}
)
