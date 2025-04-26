package main

import "fmt"

type Animal interface {
  Speak() string
}

type Dog struct {}

func (d Dog) Speak() string {
  return "Woof!"
}

func main() {
  var a Animal = Dog{}

  fmt.Println(a.Speak())
}
