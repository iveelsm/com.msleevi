package users

import (
	"api/users/handlers"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {

}

func LoginUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func NewUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	result, err := handlers.HandleUsers(params)
	if err != nil {
		w.WriteHeader(500)
		log.Error("Unable to GET /users")
		return
	}
	bytes, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(504)
		log.Error("Unable to parse data into return body")
		return
	}
	w.Write(bytes)
	return
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func SubscribeUsersHandler(w http.ResponseWriter, r *http.Request) {

}
