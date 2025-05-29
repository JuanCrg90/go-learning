package main

import (
  "fmt"
  "sync"
)

func main() {
  urls := []string{
    "https://golang.org",
    "https://httpbin.org/delay/2",
    "https://example.com",
  }

	var wg sync.WaitGroup
  ch := make(chan Result, len(urls))

  for _, url := range(urls) {
    wg.Add(1)
    go fetch(url, ch, &wg)
  }

  wg.Wait()
  close(ch)

  for res := range ch {
    fmt.Printf("%s -> %s in %v\n", res.URL, res.Status, res.Duration)
  }

  fmt.Println("All done.")
}
