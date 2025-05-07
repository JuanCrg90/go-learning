package models

import (
  "encoding/json"
  "os"
)

type Book struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
  Published bool `json:"published"`
}

type BookList struct {
  Books []Book `json:"books"`
}

func LoadBookList() (BookList, error) {
  content, err := os.ReadFile("books.json")

  if err != nil {
    return BookList{}, err
  }

  var books = BookList{}

  json.Unmarshal([]byte(content), &books)

  return books, nil
}

func SaveBookList(bookList BookList)(string, error) {
  jsonData, err := json.MarshalIndent(bookList, "", " ")

  if err != nil {
    return "", err
  }

  file, err := os.Create("books.json")

  if err != nil {
    return "", err
  }

  defer file.Close()

  _, err = file.Write(jsonData)

  if err != nil {
    return "", err
  }

  return "Book added", nil
}
