package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	product_id := query.Get("id")
	result, err := service.GetProducts(product_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func AddProducts(w http.ResponseWriter, r *http.Request) {
	var productData []byte
	_, err := r.Body.Read(productData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := service.AddProduct(productData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func UpdateProducts(w http.ResponseWriter, r *http.Request) {
	var productData []byte
	_, err := r.Body.Read(productData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := service.UpdateProduct(productData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteProducts(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	product_id := query.Get("id")

	err := service.DeleteProduct(product_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
