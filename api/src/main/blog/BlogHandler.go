package blog

import (
    "fmt"
    "net/http"
)

func HandleBlogRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hit blog endpoint")
}
