package config

import "newBackend/constants"

// app config
const (
	AppPort = 8080
	AppName = "backend"
)

// database config
const (
	DBService  = constants.MySQL
	DBHost     = "127.0.0.1"
	DBPort     = 3306
	DBUser     = "root"
	DBPassword = "deadshot21"
	DBName     = "inventory_management"
)
