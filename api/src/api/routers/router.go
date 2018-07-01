package routers

import (
	"api/blogs"
	"api/contact"
	"api/cv"
	"api/projects"
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
	projects.AddProjectsRoutes(r)
	blogs.AddBlogsRoutes(r)
	cv.AddResumeRoutes(r)
	contact.AddContactRoutes(r)
	logRoutes(r)
	return r
}

func logRoutes(r *mux.Router) {
	r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			methods = []string{}
		}
		log.WithFields(log.Fields{
			"route":   path,
			"methods": methods,
		}).Info("Route Information:")
		return nil
	})
}
