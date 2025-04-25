package main

import "fmt"

type Counter struct {
  Value int
}

func (c *Counter) Increment() {
  c.Value++
}

func main() {
  counter := Counter{ Value: 0}
  for i:=0; i< 10; i++ {
    fmt.Println(counter.Value)
    counter.Increment()
  }
}
