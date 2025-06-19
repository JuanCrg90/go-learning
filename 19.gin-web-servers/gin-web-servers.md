### ✅ **Routing & middleware (using Gin or Echo)**

#### 🔀 What you’ll learn:

* How to move beyond the standard `net/http` router
* Defining routes with parameters (e.g. `/user/:id`)
* Grouping routes
* Using middleware for things like logging, CORS, and authentication

---

### 🚀 Step 1: Choose a framework

We'll use **[Gin](https://github.com/gin-gonic/gin)** — it's fast, minimalist, and widely used in production.

---

### ✅ Install Gin:

Run this inside your project:

```bash
go get -u github.com/gin-gonic/gin
```

---

### ✅ Basic Gin Example

Create or update `main.go` with:

```go
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "Welcome to the homepage")
  })

  r.GET("/hello", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Hello World"})
  })

  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "pong"})
  })

  r.GET("/api/status", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "API is running"})
  })

  r.Run(":8080") // default is :8080
}
```
