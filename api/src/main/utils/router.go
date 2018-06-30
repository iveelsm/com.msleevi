package utils

import (
	"main/users"

	"github.com/gorilla/mux"
)

func ConstructRouter() mux.Router {
	r := mux.NewRouter()
	users.AddUsersRoutes(r)
	return r
}
