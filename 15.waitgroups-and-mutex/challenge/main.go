package main

import (
	"fmt"
	"sync"
)


func routine(r int, wg *sync.WaitGroup) {
  defer wg.Done()
  fmt.Println("Routine", r,"working")
  fmt.Println("Routine", r,"done")
}

func main()  {
  var wg sync.WaitGroup

  for _, x := range []int{1,2,3}  {
    wg.Add(1)
    go routine(x, &wg)
  }

  wg.Wait()
  fmt.Println("Main done")
}
