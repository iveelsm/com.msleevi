package main

import (
    "log"
    "net/http"

    "main/projects"
    "main/cv"
    "main/blog"
    "main/admin"
    "main/contact"

    "github.com/gorilla/mux"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
  blog.HandleBlogRequest(w, r)
}

func CVHandler(w http.ResponseWriter, r *http.Request) {
  cv.HandleCVRequest(w, r)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
  admin.HandleAdminRequest(w, r)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
  projects.HandleProjectsRequest(w, r)
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
  contact.HandleContactRequest(w, r)
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/blog", BlogHandler)
  r.HandleFunc("/projects", ProjectsHandler)
  r.HandleFunc("/admin", AdminHandler)
  r.HandleFunc("/cv", CVHandler)
  r.HandleFunc("/contact", ContactHandler)
  if err := http.ListenAndServe(":8080", r); err != nil {
    log.Fatal(err)
  }

}
