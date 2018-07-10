package main

import (
	"net/http"

	"github.com/msleevi/com.msleevi/internal/routers"
	"github.com/msleevi/com.msleevi/internal/utils"

	log "github.com/sirupsen/logrus"
)

func main() {
	utils.ConfigureLogger()
	r := routers.ConstructRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
