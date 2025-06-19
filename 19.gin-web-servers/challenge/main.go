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

  r.Run(":8080")
}
