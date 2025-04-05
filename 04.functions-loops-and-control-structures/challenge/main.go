package main

import "fmt"

func isEven(num int) bool {
  return num % 2 == 0
}

func main() {

  for i := 1 ; i <= 20 ; i++ {
    if isEven(i) {
      fmt.Println(i, "is even")
    } else {
      fmt.Println(i, "is odd")
    }
  }
}
