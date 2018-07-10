package cv

import (
	"github.com/msleevi/com.msleevi/internal/constants"
	"github.com/msleevi/com.msleevi/internal/routes/cv/handlers"

	"github.com/gorilla/mux"
)

func AddResumeRoutes(r *mux.Router) {
	r.HandleFunc("/resume", handlers.ResumeHandler).Methods(constants.GET)
}
