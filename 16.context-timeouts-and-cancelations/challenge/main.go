package main

import (
  "context"
  "fmt"
  "time"
  "sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
  defer wg.Done()

  for {
    select {
    case <-ctx.Done():
      fmt.Println("Routine cancelled:", ctx.Err())
      return
    default:
      fmt.Println("Working...")
      time.Sleep(200 * time.Millisecond)
    }
  }
}

func main()  {
  var wg sync.WaitGroup
  ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)

  defer cancel()

  wg.Add(1)
  go worker(ctx, &wg)

  wg.Wait()
  fmt.Println("Main done")
}
