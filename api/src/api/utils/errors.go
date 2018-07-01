package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func HandleError(message string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Write([]byte(message))
	log.Error(message)
	return
}
