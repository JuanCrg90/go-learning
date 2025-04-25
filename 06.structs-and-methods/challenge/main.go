/*
 Challenge: Write a program that defines a Book struct with the following fields:

Title (string)
Author (string)
Pages (int)
Create two methods:

Summary() string – returns a string like "Title by Author, 123 pages"
AddPages(extra int) – adds pages to the book (use a pointer receiver)

Then:
Create a book, print its summary, add 50 pages, and print the new summary.
*/

package main

import (
  "fmt"
  "strconv"
)

type Book struct {
  Title string
  Author string
  Pages int
}

func (b Book) Summary() string {
  strPages:= strconv.Itoa(b.Pages)
  return b.Title + " by " + b.Author +", " + strPages + " pages"
}

func (b *Book) AddPages(extra int) {
  b.Pages = b.Pages + extra
}

func main() {
  book := Book{ Title: "Go for Rubyst", Author: "JuanCrg90", Pages: 150 }

  fmt.Println(book.Summary())
  fmt.Println("Add 50 pages")
  book.AddPages(50)
  fmt.Println(book.Summary())


}
