package service

import "newBackend/connector"

func GetCustomer(id string) (any, error) {
	var fetchAll bool
	if id == "" {
		fetchAll = true
	}
	return connector.FetchCustomer(fetchAll, id)
}
