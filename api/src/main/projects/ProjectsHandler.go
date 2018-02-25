package projects

import (
    "fmt"
    "net/http"
)

func HandleProjectsRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hit Projects endpoint")
}
