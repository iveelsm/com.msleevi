package db

import (
	"github.com/go-kivik/kivik"
	log "github.com/sirupsen/logrus"
)

var connection *kivik.Client

func GetConnection() (client *kivik.Client) {
	if connection == nil {
		initializeConnection()
	}
	return connection
}

func initializeConnection() {
	client, err := constructConnection()
	if err != nil {
		log.Error("Failed to construct a database connection")
	}
	connection = client
}
