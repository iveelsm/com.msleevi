package routes

import (
	"api/utils"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Handle does
func Handle(handler func(http.ResponseWriter, *http.Request) ([]byte, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		result, code, err := handler(w, r)

		if err != nil {
			handleError(err.Error(), code, w)
		} else {
			handleSuccess(result, code, w)
		}

		log.WithFields(log.Fields{
			"location":    r.RequestURI,
			"method":      r.Method,
			"elapsedTime": utils.NanoToMillis(time.Since(start)),
		}).Info("API Call Measurements")
	}
}

func handleSuccess(bytes []byte, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Write(bytes)
}

func handleError(message string, code int, w http.ResponseWriter) {
	w.WriteHeader(code)
	w.Write([]byte(message))
	log.Error(message)
	return
}
