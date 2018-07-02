package users

import (
	"api/constants"
	"api/routes/users/handlers"

	"github.com/gorilla/mux"
)

/*
AddUsersRoutes does
*/
func AddUsersRoutes(r *mux.Router) {
	r.HandleFunc("/users/login", handlers.LoginUsersHandler).Methods(constants.GET)
	r.HandleFunc("/users/new", handlers.NewUsersHandler).Methods(constants.POST)
	r.HandleFunc("/users", handlers.GetUsersHandler).Methods(constants.GET)
	r.HandleFunc("/users/remove", handlers.RemoveUsersHandler).Methods(constants.DELETE)
	r.HandleFunc("/users/subscribe", handlers.SubscribeUsersHandler).Methods(constants.POST)
}
