package db

import (
	"github.com/go-kivik/kivik"
)

func constructConnection() (client *kivik.Client, err error) {
	client, err = kivik.New(constructContext(), "couch", "test")
	return
}
