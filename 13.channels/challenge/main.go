package main

import "fmt"

func routine(id int, ch chan string) {
  ch <- fmt.Sprintf("Routine %d done", id)
}

func main()  {
  ch := make(chan string)

  for _, x := range []int{1,2,3}  {
    go routine(x, ch)
  }

  for i := 0; i < 3; i++ {
    msg := <- ch
    fmt.Println(msg)
  }

  fmt.Println("Main function done")
}
