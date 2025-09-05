package handlers

import (
	"encoding/json"
	"net/http"
)

type HelloResponse struct {
	Message string `json:"message"`
}

// HelloHandler 返回一个 JSON API
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	resp := HelloResponse{Message: "Hello, Go!"}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(resp)
}
