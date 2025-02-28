package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func GetSuppliers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	supplier_id := query.Get("id")
	result, err := service.GetSupplier(supplier_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(result)
}

func DeleteSuppliers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	supplier_id := query.Get("id")
	err := service.DeleteSupplier(supplier_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
