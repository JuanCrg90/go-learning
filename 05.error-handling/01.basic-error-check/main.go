package main

import (
  "errors"
  "fmt"
)

func greet(name string) (string, error) {
  if name == "" {
    return "", errors.New("name cannot be empty")
  }

  return "Hello, " + name, nil
}

func main() {
  msg, err := greet("")

  if err != nil {
    fmt.Println("Error:", err)
    return
  }
  fmt.Println(msg)
}
