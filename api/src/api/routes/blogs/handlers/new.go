package handlers

import (
	"api/routes"
	"api/routes/blogs/formats"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func NewBlogPostsHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		routes.HandleError("Unable to parse body of request", 504, w)
	}

	post := formats.Blog{}
	err = json.Unmarshal(body, post)
	if err != nil {
		routes.HandleError("Unable to convert to a blog", 504, w)
	}

	result, err := handleNewPost(post)
	if err != nil {
		routes.HandleError("Unable to add new post to the database", 504, w)
	}

	w.WriteHeader(201)
	w.Write([]byte(result))
}

func handleNewPost(formats.Blog) (result string, err error) {
	return "Successful post", nil
}
