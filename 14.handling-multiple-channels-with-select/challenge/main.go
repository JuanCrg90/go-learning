package main

import (
  "fmt"
  "time"
  "math/rand"
)

func main() {
  ch1 := make(chan string)
  ch2 := make(chan string)

  go func() {
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    ch1 <- "Message from ch1"
  }()

  go func() {
    time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
    ch2 <- "Message from ch2"
  }()

  select {
  case msg := <-ch1:
    fmt.Println("Received:", msg)
  case msg := <-ch2:
    fmt.Println("Received:", msg)
  }
}
