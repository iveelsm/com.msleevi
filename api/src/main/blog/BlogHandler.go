package blog

import (
    "fmt"
    "net/http"

    "main/utils"
)

func HandleBlogRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case utils.GET_REQUEST:
      getBlogPosts(w, r)
    case utils.POST_REQUEST:
      postBlogPosts(w, r)
    case utils.DELETE_REQUEST:
      deleteBlogPosts(w, r)
    default:
      fmt.Fprintln(w, "Unsupported HTTP operation")
  }
}

func getBlogPosts(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "GET Request on Blog Endpoint")
}

func postBlogPosts(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "POST Request on Blog Endpoint")
}

func deleteBlogPosts(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "DELETE Request on Blog Endpoint")
}
