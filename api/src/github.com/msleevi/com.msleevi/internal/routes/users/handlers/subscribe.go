package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// SubscribeUser does
func SubscribeUser(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)
	result, err := subscribeUser(params)

	if err != nil {
		return nil, 504, errors.New("Unable to subscribe user")
	}

	return []byte(result), 204, nil
}

func subscribeUser(params map[string]string) (result string, err error) {
	return "successfuly subscribed user", nil
}
