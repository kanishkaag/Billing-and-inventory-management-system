package service

import (
	"newBackend/connector"
	"newBackend/models"
)

func GetInventory(shopID string) (any, error) {
	var query string
	if shopID == "" {
		query = `SELECT * FROM inventory`
	} else {
		query = `SELECT * FROM inventory WHERE shop_id = ?`
	}
	var result []models.Inventory
	err := connector.FetchInventory(query, &result, shopID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
