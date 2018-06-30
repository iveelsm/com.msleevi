package main

import (
	"log"
	"net/http"

	"main/utils"
)

func main() {
	r := utils.ConstructRouter()
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
