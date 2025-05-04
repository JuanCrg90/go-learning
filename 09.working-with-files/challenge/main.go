package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  file, err := os.Create("log.txt")

  if err != nil {
    panic(err)
  }

  _, err = file.WriteString("Log started\n")
  _, err = file.WriteString("User: Juan Carlos\n")
  _, err = file.WriteString("Status: Active\n")

  file.Close()

  file, err = os.Open("log.txt")

  if err != nil {
    panic(err)
  }

  defer file.Close()

  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
  }

  if err := scanner.Err(); err != nil {
    panic(err)
  }
}
