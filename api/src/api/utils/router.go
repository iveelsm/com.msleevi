package utils

import (
	"api/users"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

/*
ConstructRouter does
*/
func ConstructRouter() *mux.Router {
	r := mux.NewRouter()
	users.AddUsersRoutes(r)
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		t, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		log.Info(t)
		return nil
	})
	return r
}
