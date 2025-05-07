package main

import (
	"challenge/handlers"
	"net/http"
)

func main() {
  http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
      handlers.ListBooks(w, r)
    } else if r.Method == http.MethodPost {
      handlers.AddBook(w, r)
    } else {
      http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
  })
  http.HandleFunc("/books/{id}", handlers.GetBook)


  http.ListenAndServe(":8080", nil)
}
