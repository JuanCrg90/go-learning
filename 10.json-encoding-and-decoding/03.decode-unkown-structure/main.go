package main

import (
  "encoding/json"
  "fmt"
)

func main() {
  input := `{"name": "Laura", "age": 30}`
  var m map[string]interface{}

  json.Unmarshal([]byte(input), &m)

  for k,v := range m {
    fmt.Printf("%s: %v\n", k, v)
  }
}
