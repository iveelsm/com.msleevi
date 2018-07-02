package handlers

import (
	"api/routes"
	"api/routes/blogs/formats"
	"net/http"

	"github.com/gorilla/mux"
)

func BlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	_, err := handleBlogs(params)
	if err != nil {
		routes.HandleError("Unable to find blog posts", 500, w)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("foo"))
}

func handleBlogs(params map[string]string) (results []formats.Blog, err error) {
	return []formats.Blog{}, nil
}
