package main

import "fmt"

func main() {
  numbers := []int{1, 2, 3}

  for index, value := range numbers {
    fmt.Println(index, value)
  }
}
