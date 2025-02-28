package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"newBackend/app"
	"newBackend/config"
	"newBackend/database"
	"newBackend/models"
)

func main() {
	logrus.Infof("Server is listening on port %v...", config.AppPort)
	logrus.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppPort), app.Router))
}

func init() {
	var err error
	models.DBClient, err = database.GetDBClient()
	if err != nil {
		logrus.Fatal("failed to establish connection with database service; ", err)
	}

	app.InitHandlers()
}
