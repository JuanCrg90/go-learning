/*
 Challenge: Write a Go program that:

Declares a const named BaseSalary with a value of 50000.
Declares a var named bonus initialized to 5000.
Calculates the totalSalary by adding BaseSalary + bonus.
Converts the totalSalary to a string and prints a message like:
"The total salary is 55000"
*/

package main

import (
  "fmt"
  "strconv"
)

const BaseSalary = 50000

func main()  {
  bonus := 5000

  totalSalary := BaseSalary + bonus

  str := strconv.Itoa(totalSalary)

  fmt.Println("The total salary is", str)
}
