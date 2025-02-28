package service

import (
	"newBackend/connector"
	"newBackend/models"
)

func RecordPayment(id string) (any, error) {
	var query string
	if id == "" {
		// Fetch all payments if id is empty
		query = `SELECT * FROM payment`
	} else {
		// Fetch a single payment if id is provided
		query = `SELECT * FROM payment WHERE id = ?`
	}

	var result []models.Payment
	err := connector.FetchPayment(query, &result, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SavePayment(billID string, status string, medium string) (any, error) {
	return connector.WritePayment(billID, status, medium)
}
