package main

import "fmt"

func increment(x *int) {
  *x++
}

func main() {
  num := 5

  fmt.Println(num)
  increment(&num)
  fmt.Println(num)
}
