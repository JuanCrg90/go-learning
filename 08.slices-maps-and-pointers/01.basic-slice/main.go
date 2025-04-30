package main

import "fmt"

func main() {
  fruits := []string{"apple", "banana", "cherry"}

  for _, fruit := range fruits {
    fmt.Println(fruit)
  }
}
