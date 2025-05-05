package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
  Name string
  Email string
}

func main() {
  input := `{"name": "Juan", "email":"juan@example.com"}`
  var u User
  json.Unmarshal([]byte(input), &u)
  fmt.Println(u)
}
