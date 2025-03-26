package main

import "fmt"

func divide(a, b int) (int, error) {
  if b == 0 {
    return 0, fmt.Errorf("Division by zero")
  }

  return a/b, nil
}

func main()  {
  fmt.Println(divide(8,2))
  fmt.Println(divide(8,0))
}
