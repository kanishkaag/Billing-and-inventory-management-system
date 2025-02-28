package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func HandleGetBills(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	bill_id := query.Get("id")

	result, err := service.GetBill(bill_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func HandleSaveBill(w http.ResponseWriter, r *http.Request) {
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

	result, err := service.SaveBill(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(result)
}
