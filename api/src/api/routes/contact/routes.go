package contact

import (
	"api/constants"
	"api/routes/contact/handlers"

	"github.com/gorilla/mux"
)

/*
AddContactRoutes does
*/
func AddContactRoutes(r *mux.Router) {
	r.HandleFunc("/contact", handlers.ContactHandler).Methods(constants.POST)
}
