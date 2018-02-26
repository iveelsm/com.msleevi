package cv

import (
    "fmt"
    "net/http"

    "main/utils"
)

func HandleCVRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case utils.GET_REQUEST:
      getCV(w, r)
    case utils.POST_REQUEST:
      postCV(w, r)
    case utils.DELETE_REQUEST:
      deleteCV(w, r)
    default:
      fmt.Fprintln(w, "Unsupported HTTP operation")
  }
}

func getCV(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "GET Request on CV Endpoint")
}

func postCV(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "POST Request on CV Endpoint")
}

func deleteCV(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "DELETE Request on CV Endpoint")
}
