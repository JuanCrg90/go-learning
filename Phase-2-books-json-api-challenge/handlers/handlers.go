package handlers

import (
  "encoding/json"
  "net/http"
  "strconv"
  "challenge/models"
)

func ListBooks(w http.ResponseWriter, r *http.Request) {
  bookList, err := models.LoadBookList()
  if err != nil {
    http.Error(w, "failed to load books", http.StatusInternalServerError)
    return
  }

  jsonData, err := json.MarshalIndent(bookList,"", " ")

  if err != nil {
    http.Error(w, "failed to decode json", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonData)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
  bookList, err := models.LoadBookList()
  if err != nil {
    http.Error(w, "failed to load books", http.StatusInternalServerError)
    return
  }

  idStr := r.PathValue("id")
  id, err := strconv.Atoi(idStr)

  if err != nil || id < 1 || id > len(bookList.Books) {
    http.Error(w, "book not found", http.StatusNotFound)
    return
  }

  book := bookList.Books[id - 1]

  jsonData, err := json.MarshalIndent(book,"", " ")

  if err != nil {
    http.Error(w, "could not encode the book", http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonData)
}


func AddBook(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
  }

  var newBook models.Book

  if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
    http.Error(w, "Invalid JSON", http.StatusBadRequest)
    return
  }

  bookList, err := models.LoadBookList()
  if err != nil {
    http.Error(w, "failed to load books", http.StatusInternalServerError)
    return
  }

  maxID := 0

  for _, book := range bookList.Books {
    if book.ID > maxID {
      maxID = book.ID
    }
  }

  newBook.ID = maxID + 1

  bookList.Books = append(bookList.Books, newBook)

  msg, err := models.SaveBookList(bookList)

  if err != nil {
    http.Error(w, "could not save book", http.StatusInternalServerError)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  json.NewEncoder(w).Encode(map[string]string{"message": msg})
}
