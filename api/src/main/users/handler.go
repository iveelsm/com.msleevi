package users

import (
	"net/http"

	"github.com/gorilla/mux"
)

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {

}

func LoginUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func NewUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
}

func DeleteUsersHandler(w http.ResponseWriter, r *http.Request) {

}

func SubscribeUsersHandler(w http.ResponseWriter, r *http.Request) {

}
