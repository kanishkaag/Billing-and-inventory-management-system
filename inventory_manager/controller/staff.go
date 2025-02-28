package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func GetStaff(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	Staff_id := query.Get("id")
	result, err := service.GetStaff(Staff_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(result)
}
