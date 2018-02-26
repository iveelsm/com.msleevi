package Contact

import (
    "fmt"
    "net/http"

    "main/utils"
)

func HandleContactRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case utils.GET_REQUEST:
      getContact(w, r)
    case utils.POST_REQUEST:
      postContact(w, r)
    case utils.DELETE_REQUEST:
      deleteContact(w, r)
    default:
      fmt.Fprintln(w, "Unsupported HTTP operation")
  }
}

func getContact(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "GET Request on Contact Endpoint")
}

func postContact(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "POST Request on Contact Endpoint")
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "DELETE Request on Contact Endpoint")
}
