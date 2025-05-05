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
  u := User{Name: "Juan", Email: "juan@example.com"}
  jsonData, _ := json.MarshalIndent(u,"", " ")
  fmt.Println(string(jsonData))
}
