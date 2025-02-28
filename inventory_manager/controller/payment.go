package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func RecordPayment(w http.ResponseWriter, r *http.Request) {
	var paymentData []byte
	_, err := r.Body.Read(paymentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := service.RecordPayment(paymentData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}
