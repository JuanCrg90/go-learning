package main

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func main() {
  http.HandleFunc("/", homeHandler)
  http.HandleFunc("/hello", helloHandler)
  http.HandleFunc("/ping", pingHandler)
  http.HandleFunc("/api/status", statusHandler)

  fmt.Println("Server Listening on http://localhost:8080")
  http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome to the homepage")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Hello World")
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Pong")
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  response := map[string]string {
    "status": "ok",
    "message": "API is running",
  }

  json.NewEncoder(w).Encode(response)
}
