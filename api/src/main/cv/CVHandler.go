package cv

import (
    "fmt"
    "net/http"
)

func HandleCVRequest(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "hit CV endpoint")
}
