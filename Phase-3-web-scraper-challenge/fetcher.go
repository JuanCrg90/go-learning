package main

import (
  "net/http"
  "time"
  "sync"
)

func fetch(url string, ch chan<- Result, wg *sync.WaitGroup) {
  defer wg.Done()

  start := time.Now()
  resp, err := http.Get(url)
  duration := time.Since(start)

  if err != nil {
    ch <- Result{URL: url, Duration: duration, Status: "Error" }
  }

  defer resp.Body.Close()

  ch <- Result{URL: url, Duration: duration, Status: resp.Status }
}
