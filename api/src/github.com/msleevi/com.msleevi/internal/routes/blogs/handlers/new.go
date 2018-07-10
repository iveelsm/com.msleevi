package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/msleevi/com.msleevi/internal/routes/blogs/formats"
)

func NewBlogPostsHandler(w http.ResponseWriter, r *http.Request) (ret []byte, code int, err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, 504, errors.New("Unable to parse body of request")
	}

	post := formats.Blog{}
	err = json.Unmarshal(body, post)
	if err != nil {
		return nil, 504, errors.New("Unable to convert to a blog")
	}

	result, err := handleNewPost(post)
	if err != nil {
		return nil, 500, errors.New("Uanble to add new post to the database")
	}

	return []byte(result), 201, nil
}

func handleNewPost(formats.Blog) (result string, err error) {
	return "Successful post", nil
}
