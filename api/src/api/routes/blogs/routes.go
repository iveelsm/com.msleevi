package blogs

import (
	"api/constants"
	"api/routes"
	"api/routes/blogs/handlers"

	"github.com/gorilla/mux"
)

/*
AddBlogsRoutes does
*/
func AddBlogsRoutes(r *mux.Router) {
	r.HandleFunc("/blogs", routes.Handle(handlers.BlogPostsHandler)).Methods(constants.GET)
	r.HandleFunc("/blogs/new", routes.Handle(handlers.NewBlogPostsHandler)).Methods(constants.POST)
	r.HandleFunc("/blogs/remove", routes.Handle(handlers.RemoveBlogPostsHandler)).Methods(constants.DELETE)
}
