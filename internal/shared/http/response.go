package http

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Data 	any `json:"data"`
	Error 	any `json:"error"`
}

func Success(w http.ResponseWriter, status int, data any){
	response := Response{
		Data: data,
		Error: nil,
	}

	writeJSON(w, status, response)
}

func Error(w http.ResponseWriter, status int, err error){
	response := Response{
		Data: nil,
		Error: err.Error(),
	}

	writeJSON(w, status, response)

}

func writeJSON(w http.ResponseWriter, status int, resp Response){
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(resp)
}