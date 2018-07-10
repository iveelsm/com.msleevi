package routers

import (
	"github.com/msleevi/com.msleevi/internal/routes/contact"
	"github.com/msleevi/com.msleevi/internal/routes/cv"
	"github.com/msleevi/com.msleevi/internal/routes/projects"
	"github.com/msleevi/com.msleevi/internal/routes/users"

	"github.com/msleevi/com.msleevi/internal/routes/blogs"

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
