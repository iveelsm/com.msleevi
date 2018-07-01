package projects

import (
	"api/constants"
	"api/projects/handlers"

	"github.com/gorilla/mux"
)

/*
AddProjectsRoutes does
*/
func AddProjectsRoutes(r *mux.Router) {
	r.HandleFunc("/projects/new", handlers.NewProjectsHandler).Methods(constants.POST)
	r.HandleFunc("/projects", handlers.ProjectsHandler).Methods(constants.GET)
	r.HandleFunc("/projects/remove", handlers.RemoveProjectsHandler).Methods(constants.DELETE)
}
