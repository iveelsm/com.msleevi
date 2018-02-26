package admin

import (
    "fmt"
    "net/http"

    "main/utils"
)

func HandleAdminRequest(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case utils.GET_REQUEST:
      getAdminStatus(w, r)
    case utils.POST_REQUEST:
      postAdminStatus(w, r)
    case utils.DELETE_REQUEST:
      deleteAdminStatus(w, r)
    default:
      fmt.Fprintln(w, "Unsupported HTTP operation")
  }
}

func getAdminStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "GET Request on Admin Endpoint")
}

func postAdminStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "POST Request on Admin Endpoint")
}

func deleteAdminStatus(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "DELETE Request on Admin Endpoint")
}
