package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func HandleListProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	product_id := query.Get("name")
	result, err := service.GetProducts(product_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}
