package database

import (
	"database/sql"
	"errors"
	"newBackend/config"
	"newBackend/constants"
)

type Client interface {
	WriteOne(query string, args ...interface{}) (sql.Result, error)
	WriteMany(query string, records []interface{}) error
	ReadOne(dest interface{}, query string, args ...interface{}) error
	ReadMany(dest interface{}, query string, args ...interface{}) error
	UpdateOne(query string, args ...interface{}) (sql.Result, error)
	UpdateMany(string, []interface{}) error
	DeleteOne(query string, args ...interface{}) (sql.Result, error)
	DeleteMany(query string, args ...interface{}) (sql.Result, error)
}

type connectionConfig struct {
	host     string
	port     int
	username string
	password string
	database string
}

func GetDBClient() (Client, error) {
	connConfig := connectionConfig{
		host:     config.DBHost,
		port:     config.DBPort,
		username: config.DBUser,
		password: config.DBPassword,
		database: config.DBName,
	}

	switch config.DBService {

	case constants.MySQL:
		return getMySQLClient(connConfig)

	default:
		return nil, errors.New("invalid DB service: " + config.DBService)
	}
}
