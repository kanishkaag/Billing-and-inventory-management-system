package service

import "newBackend/connector"

func GetStaff(id string) (any, error) {
	var fetchAll bool
	if id == "" {
		fetchAll = true
	}
	return connector.FetchStaff(fetchAll, id)
}
