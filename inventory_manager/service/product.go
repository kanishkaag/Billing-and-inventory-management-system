package service

import (
	"errors"
	"newBackend/connector"
)

// GetProducts fetches all products if id is empty, otherwise fetches a specific product by id.
func GetProducts(id string) (any, error) {
	// Determine if we need to fetch all products or a specific one based on the id.
	var fetchAll bool
	if id == "" {
		fetchAll = true
	}
	return connector.FetchProduct(fetchAll, id)
}

// AddProduct inserts a new product into the database.
func AddProduct(data []byte) (any, error) {
	// Call the connector function to insert a new product.
	if data == nil || len(data) == 0 {
		return nil, errors.New("invalid data for product")
	}
	return connector.InsertProduct(data)
}

// UpdateProduct modifies an existing product's details in the database.
func UpdateProduct(data []byte) (any, error) {
	// Call the connector function to update the product.
	if data == nil || len(data) == 0 {
		return nil, errors.New("invalid data for updating product")
	}
	return connector.ModifyProduct(data)
}

// DeleteProduct deletes a product from the database by its id.
func DeleteProduct(id string) error {
	// Ensure that an id is provided before trying to delete the product.
	if id == "" {
		return errors.New("product id must be provided to delete")
	}
	return connector.RemoveProduct(id)
}
