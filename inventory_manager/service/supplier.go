package service

import (
	"errors"
	"newBackend/connector"
	"newBackend/models"
)

// GetSupplier fetches a supplier based on the provided id, or fetches all suppliers if id is empty.
func GetSupplier(id string) (any, error) {
	var query string
	// If id is empty, fetch all suppliers, else fetch the supplier with the given id
	if id == "" {
		query = `SELECT * FROM supplier`
	} else {
		query = `SELECT * FROM supplier WHERE id = ?`
	}

	var result []models.Supplier
	// Call the connector function to fetch the supplier(s)
	err := connector.FetchSupplier(query, &result, id)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteSupplier removes a supplier based on the given id.
func DeleteSupplier(id string) error {
	if id == "" {
		return errors.New("product id must be provided to delete")
	}
	return connector.RemoveSupplier(id)
}
