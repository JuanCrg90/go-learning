## 🔧 **Project: Book Library API**

### **🎯 Overview**

Build a small REST-style JSON API that lets you:

* List all books
* Add a new book
* Get details for a single book by ID
* (Optional) Delete a book

Data is stored in a **JSON file** (`books.json`) and loaded/saved on each request.

---

### **📦 Suggested File Structure**

```
challenge/
├── go.mod
├── main.go
├── handlers/
│   └── handlers.go
├── models/
│   └── models.go
└── books_test.go
```

---

### **📚 Sample Book Struct**

```go
type Book struct {
    ID       int    `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Published bool  `json:"published"`
}
```

---

### **📌 Required Endpoints (using `net/http`)**

| Method | Path          | Description                |
| ------ | ------------- | -------------------------- |
| GET    | `/books`      | List all books             |
| GET    | `/books/{id}` | Get a book by ID           |
| POST   | `/books`      | Add a new book (from JSON) |

(Optional: DELETE `/books/{id}`)

---

### **🧠 Key Features to Implement**

* Use `encoding/json` for request/response.
* Use a slice `[]Book` as in-memory storage.
* Read/write the file using `os.ReadFile` / `os.WriteFile`.
* Implement basic routing with `http.HandleFunc`.
* (Optional) Use an `interface{}` for generic error response formats.

---

### **✅ How This Reinforces Phase 2 Skills**

| Skill         | Application                                     |
| ------------- | ----------------------------------------------- |
| Structs       | Define and model `Book` data                    |
| Interfaces    | Implement basic abstraction for storage methods |
| Maps & Slices | Store book records dynamically                  |
| File I/O      | Load and persist data in JSON                   |
| JSON Encoding | API request and response handling               |
| Testing       | Unit tests for handlers or storage logic        |

