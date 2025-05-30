## ✅ `net/http` for building web servers in Go

---

### 🔧 What you'll learn in this section:

* Start a basic HTTP server
* Define handlers for different routes (e.g., `/`, `/hello`)
* Understand how requests and responses work
* Return HTML or JSON manually

---

## 📘 1. Basic Web Server Structure

Here’s a minimal Go web server using `net/http`:

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page")
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Juan Carlos 👋")
}
```

> Run this with `go run .` and visit:
>
> * [http://localhost:8080/](http://localhost:8080/)
> * [http://localhost:8080/hello](http://localhost:8080/hello)

---

## 🧠 Key Concepts

| Concept                             | Description                                                        |
| ----------------------------------- | ------------------------------------------------------------------ |
| `http.HandleFunc`                   | Registers a route with a handler function                          |
| `http.ResponseWriter`               | Where you write the response to the client                         |
| `*http.Request`                     | Contains info about the request (method, headers, URL, body, etc.) |
| `http.ListenAndServe(":8080", nil)` | Starts the server on port 8080                                     |

---

### ✅ Mini Challenge:

Can you add a new route called `/ping` that returns `"pong"`?


### 🔄 Next: Returning JSON Responses

Let’s take it a step further by returning JSON instead of plain text. This is essential for building real APIs.

Here’s how you can do it:

#### 🧪 Add a new route: `/api/status`

```go
func statusHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  fmt.Fprint(w, `{"status": "ok", "message": "API is running"}`)
}
```

And register it in `main()`:

```go
http.HandleFunc("/api/status", statusHandler)
```

### ✅ Optional: Use `json.Marshal`

For real APIs, it's better to build your JSON from a struct:

```go
import "encoding/json"

func statusHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)

  response := map[string]string{
    "status":  "ok",
    "message": "API is running",
  }

  json.NewEncoder(w).Encode(response)
}
```

This ensures the output is always valid JSON — especially helpful as your responses grow more complex.

Let me know if you want to move on to **handling query params or JSON requests next.**
---
