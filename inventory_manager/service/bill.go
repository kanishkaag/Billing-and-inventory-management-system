package service

import (
	"newBackend/connector"
	"newBackend/models"
)

func GetBill(id string) (any, error) {
	var query string
	var result []models.Bill

	if id != "" {
		query = `SELECT * FROM bill WHERE id = ?`
		err := connector.FetchBill(query, &result, id)
		if err != nil {
			return nil, err
		}
	} else {
		query = `SELECT * FROM bill`
		err := connector.FetchBill(query, &result)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func SaveBill(body []byte) (any, error) {
	return connector.WriteBill(body)
}
