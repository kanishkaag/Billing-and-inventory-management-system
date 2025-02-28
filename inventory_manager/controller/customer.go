package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func HandleGetCustomers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	customer_id := query.Get("id")
	result, err := service.GetCustomer(customer_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(result)
}

func HandleAddCustomer(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var body []byte
	_, err = r.Body.Read(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := service.AddCustomer(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}
