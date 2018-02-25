package admin

import (
    "fmt"
    "net/http"
)

func HandleAdminRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hit Admin endpoint")
}
