package users

import (
	"api/constants"
	"api/routes"
	"api/routes/users/handlers"

	"github.com/gorilla/mux"
)

/*
AddUsersRoutes does
*/
func AddUsersRoutes(r *mux.Router) {
	r.HandleFunc("/users/login", routes.Handle(handlers.LoginUser)).Methods(constants.GET)
	r.HandleFunc("/users/new", routes.Handle(handlers.NewUser)).Methods(constants.POST)
	r.HandleFunc("/users", routes.Handle(handlers.GetUsers)).Methods(constants.GET)
	r.HandleFunc("/users/{id}", routes.Handle(handlers.GetUser)).Methods(constants.GET)
	r.HandleFunc("/users/remove/{id}", routes.Handle(handlers.RemoveUser)).Methods(constants.DELETE)
	r.HandleFunc("/users/subscribe/{id}", routes.Handle(handlers.SubscribeUser)).Methods(constants.POST)
}
