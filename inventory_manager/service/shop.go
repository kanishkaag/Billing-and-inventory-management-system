package service

import "newBackend/connector"

func GetShop(id string) (any, error) {
	var fetchAll bool
	if id == "" {
		fetchAll = true
	}
	return connector.FetchShop(fetchAll, id)
}
