package handlers

import (
	"api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func SubscribeUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := subscribeUser(params)

	if err != nil {
		routes.HandleError("Unable to subscribe user", 504, w)
		return
	}

	w.WriteHeader(204)
	w.Write([]byte(result))
}

func subscribeUser(params map[string]string) (result string, err error) {
	return "successfuly subscribed user", nil
}
