package utils

import (
	"main/users"

	"github.com/gorilla/mux"
)

/*
ConstructRouter does
*/
func ConstructRouter() mux.Router {
	r := mux.NewRouter()
	users.AddUsersRoutes(r)
	return r
}
