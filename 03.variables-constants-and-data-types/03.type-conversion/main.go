package main

import (
  "fmt"
  "strconv"
)

func main() {
  var num int = 42
  str := strconv.Itoa(num) // Convert int to String

  fmt.Println("String:", str)
}
