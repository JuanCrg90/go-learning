/*
Challenge: Write a function loadUser(id int) (string, error) that:

Returns "User 1" if id == 1
Returns an error saying "user not found" for any other id
Then in main(), call this function with id 2, and handle the error by printing:

Failed to load user: user not found
*/
package main

import (
  "errors"
  "fmt"
)

func loadUser(id int) (string, error) {
  if id == 1 {
    return "User 1", nil
  }

  return "", errors.New("user not found")
}

func main() {
  msg, err := loadUser(99)

  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(msg)
}
