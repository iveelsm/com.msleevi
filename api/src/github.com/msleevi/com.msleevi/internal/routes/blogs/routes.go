package blogs

import (
	"github.com/msleevi/com.msleevi/internal/constants"
	"github.com/msleevi/com.msleevi/internal/routes"
	"github.com/msleevi/com.msleevi/internal/routes/blogs/handlers"

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
