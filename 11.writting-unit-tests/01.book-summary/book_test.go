package book

import "testing"

func TestSummary(t *testing.T) {
  b := Book{"The Hobbit", "J.R.R Tolkien"}
  want := "The Hobbit by J.R.R Tolkien"
  got := b.Summary()

  if got != want {
    t.Errorf("Summary() = %q; want %q", got, want)
  }
}
