/*
Write a table-driven test to verify the Sum() function works for:
1 + 1 = 2
5 + 7 = 12
-3 + 3 = 0
*/
package calculator

import "testing"

func TestSum(t *testing.T) {
  tests := []struct {
    a, b int
    want int
  }{
    {1, 1, 2},
    {5, 7, 12},
    {-3, 3, 0},
  }

  for _, tt := range tests {
    got := Sum(tt.a, tt.b)

    if got != tt.want {
      t.Errorf("Sum(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
    }
  }
}
