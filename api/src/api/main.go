package main

import (
	"api/routers"
	"api/utils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func main() {
	utils.ConfigureLogger()
	r := routers.ConstructRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
