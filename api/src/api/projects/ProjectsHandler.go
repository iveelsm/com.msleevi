package projects

import (
    "fmt"
    "net/http"

    "main/utils"
)

func HandleProjectsRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case utils.GET_REQUEST:
      getProjects(w, r)
    case utils.POST_REQUEST:
      postProjects(w, r)
    case utils.DELETE_REQUEST:
      deleteProjects(w, r)
    default:
      fmt.Fprintln(w, "Unsupported HTTP operation")
  }
}

func getProjects(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "GET Request on Projects Endpoint")
}

func postProjects(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "POST Request on Projects Endpoint")
}

func deleteProjects(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "DELETE Request on Projects Endpoint")
}
