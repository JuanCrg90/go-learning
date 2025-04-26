package main

import "fmt"

var val interface{} = 42

func main()  {
  num, ok := val.(int)

  if ok {
    fmt.Println("It's an int:", num)
  } else {
    fmt.Println("Not an int")
  }
}
