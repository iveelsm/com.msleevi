package users

import (
	"api/constants"

	"github.com/gorilla/mux"
)

/*
AddUsersRoutes does
*/
func AddUsersRoutes(r *mux.Router) {
	r.HandleFunc("/users/login", LoginUsersHandler).Methods(constants.GET)
	r.HandleFunc("/users/new", NewUsersHandler).Methods(constants.POST)
	r.HandleFunc("/users", GetUsersHandler).Methods(constants.GET)
	r.HandleFunc("/users/remove", DeleteUsersHandler).Methods(constants.DELETE)
	r.HandleFunc("/users/subscribe", SubscribeUsersHandler).Methods(constants.POST)
}
