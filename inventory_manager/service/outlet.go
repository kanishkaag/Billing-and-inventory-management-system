package service

import (
	"newBackend/connector"
	"newBackend/models"
)

func GetOutlet(shopID string) (any, error) {
	var query string
	if shopID == "" {
		query = `SELECT * FROM outlet`
	} else {
		query = `SELECT * FROM outlet WHERE shop_id = ?`
	}
	var result []models.Outlet
	err := connector.FetchOutlet(query, &result, shopID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
