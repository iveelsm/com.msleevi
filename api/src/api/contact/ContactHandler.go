package contact

import (
    "encoding/json"
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
  var ret Contact
  decoder := json.NewDecoder(r.Body)
  err := decoder.Decode(&ret)
  if err != nil {
    fmt.Println(err)
  }
  fmt.Printf("%+v\n", ret)
  fmt.Fprintln(w, "POST Request on Contact Endpoint")
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "DELETE Request on Contact Endpoint")
}
