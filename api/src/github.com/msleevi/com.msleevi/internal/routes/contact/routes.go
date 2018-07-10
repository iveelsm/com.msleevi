package contact

import (
	"github.com/msleevi/com.msleevi/internal/constants"
	"github.com/msleevi/com.msleevi/internal/routes/contact/handlers"

	"github.com/gorilla/mux"
)

/*
AddContactRoutes does
*/
func AddContactRoutes(r *mux.Router) {
	r.HandleFunc("/contact", handlers.ContactHandler).Methods(constants.POST)
}
