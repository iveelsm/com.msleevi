package blogs

import (
	"api/blogs/handlers"
	"api/constants"

	"github.com/gorilla/mux"
)

/*
AddBlogsRoutes does
*/
func AddBlogsRoutes(r *mux.Router) {
	r.HandleFunc("/blogs", handlers.BlogPostsHandler).Methods(constants.GET)
	r.HandleFunc("/blogs/new", handlers.NewBlogPostsHandler).Methods(constants.POST)
	r.HandleFunc("/blogs/remove", handlers.RemoveBlogPostsHandler).Methods(constants.DELETE)
}
