/*
Challenge: Write a program that:

Defines a Book struct with:
Title (string)
Author (string)
Published (bool)
Creates a Book, encodes it as JSON, and writes it to book.json
Reads book.json, decodes it back into a struct, and prints the book's title and author.
*/
package main

import (
  "encoding/json"
  "fmt"
  "os"
)

type Book struct {
  Title string `json:"title"`
  Author string `json:"author"`
  Published bool `json:"published"`
}

func main() {
  book := Book{Title: "De la tierra a la luna", Author: "Julio Verne", Published: true}

  jsonData, _ := json.MarshalIndent(book, "", " ")

  file, err := os.Create("book.json")

  if err != nil {
    panic(err)
  }

  _, err = file.Write(jsonData)
  if err != nil {
    panic(err)
  }

  file.Close()

  content, err := os.ReadFile("book.json")
  if err != nil {
    panic(err)
  }

  var decodedBook Book

  json.Unmarshal(content, &decodedBook)
  fmt.Println(decodedBook.Title, "-", decodedBook.Author, "Published:", decodedBook.Published)
}
