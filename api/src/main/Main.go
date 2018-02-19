package main

import (
    "fmt"
    "html"
    "log"
    "net/http"
)

var session *http.Server
var router Router

func http.Server constructSession() {
  return &http.Server{
        Addr:           port,
        Handler:        &r,
        ReadTimeout:    10 * time.Second,
        WriteTimeout:   10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
}


func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))

}
