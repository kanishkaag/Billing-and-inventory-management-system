package controller

import (
	"encoding/json"
	"net/http"
	"newBackend/service"
)

func GetShops(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	shop_id := query.Get("id")
	result, err := service.GetShop(shop_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(result)
}
