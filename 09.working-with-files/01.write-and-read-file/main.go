package main

import (
	"fmt"
	"os"
)

func main() {
  err := os.WriteFile("hello.txt", []byte("Hello, Juan Carlos!"), 0644)
  if err != nil {
    panic(err)
  }

  data, err := os.ReadFile("hello.txt")
  if err != nil {
    panic(err)
  }

  fmt.Println(string(data))

}
