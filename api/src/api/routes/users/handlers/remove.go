package handlers

import (
	"api/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func RemoveUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := removeUser(params)

	if err != nil {
		routes.HandleError("Unable to delete user", 504, w)
	}

	w.WriteHeader(204)
	w.Write([]byte(result))
}

func removeUser(params map[string]string) (result string, err error) {
	return "deletion succesful", nil
}
