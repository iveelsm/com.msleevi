package cv

import (
	"api/constants"
	"api/cv/handlers"

	"github.com/gorilla/mux"
)

func AddResumeRoutes(r *mux.Router) {
	r.HandleFunc("/resume", handlers.ResumeHandler).Methods(constants.GET)
}
