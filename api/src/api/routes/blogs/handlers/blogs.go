package handlers

import (
	"api/routes/blogs/formats"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

func BlogPostsHandler(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	params := mux.Vars(r)

	results, err := handleBlogs(params)
	if err != nil {
		return nil, 500, errors.New("Unable to find blog posts")
	}

	ret, err = json.Marshal(results)
	if err != nil {
		return nil, 504, errors.New("Unable to parse data into return body")
	}

	return ret, 200, nil
}

func handleBlogs(params map[string]string) (results []formats.Blog, err error) {
	return []formats.Blog{}, nil
}
