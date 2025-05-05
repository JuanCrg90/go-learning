## **📌 Lesson 10: JSON Encoding & Decoding in Go (`encoding/json`)**

### **1️⃣ Introduction**

JSON is the standard format for data exchange between web servers, APIs, and frontend apps. Go provides a powerful and easy-to-use `encoding/json` package to:

* Encode (marshal) structs into JSON
* Decode (unmarshal) JSON into structs
* Handle nested, optional, and dynamic fields

---

### **2️⃣ Core Concepts**

---

#### 🔹 Encoding (Go → JSON)

Use `json.Marshal` to convert a struct or map to JSON:

```go
type User struct {
    Name  string
    Email string
}

u := User{"Juan", "juan@example.com"}
data, err := json.Marshal(u)
fmt.Println(string(data))  // {"Name":"Juan","Email":"juan@example.com"}
```

✅ Use `json.MarshalIndent()` for pretty-printed JSON.

---

#### 🔹 Decoding (JSON → Go)

Use `json.Unmarshal` to parse JSON into a Go struct:

```go
jsonData := `{"Name": "Juan", "Email": "juan@example.com"}`
var u User
err := json.Unmarshal([]byte(jsonData), &u)
fmt.Println(u.Name)  // Juan
```

---

#### 🔹 JSON Tags

Control JSON keys and behavior with struct tags:

```go
type User struct {
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}
```

* `json:"name"` → maps JSON key `"name"` to the field
* `omitempty` → skips empty fields during encoding

---

#### 🔹 Decoding into a `map[string]interface{}`

Useful for dynamic or unknown structures:

```go
var data map[string]interface{}
json.Unmarshal([]byte(jsonData), &data)
```

Then use type assertions:

```go
name := data["name"].(string)
```

---

### **3️⃣ Code Examples**

#### Example 1: Marshal a struct

```go
u := User{Name: "Juan", Email: "juan@example.com"}
jsonData, _ := json.MarshalIndent(u, "", "  ")
fmt.Println(string(jsonData))
```

#### Example 2: Unmarshal into struct

```go
input := `{"name": "Carlos", "email": "carlos@example.com"}`
var u User
json.Unmarshal([]byte(input), &u)
fmt.Println(u)
```

#### Example 3: Decode unknown structure

```go
input := `{"name": "Laura", "age": 30}`
var m map[string]interface{}
json.Unmarshal([]byte(input), &m)

for k, v := range m {
    fmt.Printf("%s: %v\n", k, v)
}
```

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Use **struct tags** for clean key mapping and JSON compatibility.
- ✅ Always check `err` from `json.Marshal/Unmarshal`.
- ✅ Use `omitempty` to avoid empty values in output.
- ❌ Don’t forget `&` when unmarshaling into a struct (`&target`).
- ❌ Be careful with `interface{}` and type assertions — use checks.

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:

1. Defines a `Book` struct with:

   * `Title` (string)
   * `Author` (string)
   * `Published` (bool)
2. Creates a `Book`, encodes it as JSON, and writes it to `book.json`
3. Reads `book.json`, decodes it back into a struct, and prints the book's title and author.

---

### **6️⃣ Further Reading & Resources**

- 📖 [Go by Example – JSON](https://gobyexample.com/json)
- 📖 [`encoding/json` docs](https://pkg.go.dev/encoding/json)
- 📖 [Effective Go – JSON and Tags](https://golang.org/doc/effective_go.html#json)
