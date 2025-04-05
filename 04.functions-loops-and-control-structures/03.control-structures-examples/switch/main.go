package main

import "fmt"

func main()  {
  day := "Tuesday"

  switch day {
  case "Monday":
    fmt.Println("Start of Week")
  case "Friday":
    fmt.Println("End of week")
  default:
    fmt.Println("Midweek")
  }
}
