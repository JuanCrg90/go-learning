/*
Challenge: Write a small testable Go program in two files:

calculator.go

Implement a Sum(a, b int) int function.
calculator_test.go

Write a table-driven test to verify the Sum() function works for:
1 + 1 = 2
5 + 7 = 12
-3 + 3 = 0
Run it with go test.
*/

package calculator

func Sum(a, b int) int {
  return a + b
}
