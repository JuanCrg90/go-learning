## **📌 Lesson 5: Go’s Unique Approach to Error Handling**

### **1️⃣ Introduction**
Unlike many languages that use exceptions (like Ruby or JavaScript),
Go handles errors **explicitly**.
This approach leads to **more predictable, readable, and robust code**,
though it might feel verbose at first.

---

### **2️⃣ Core Concepts**

#### 🔹 The `error` Type
- `error` is a built-in Go interface:
```go
type error interface {
    Error() string
}
```

- Common usage pattern:
```go
result, err := someFunction()
if err != nil {
    // handle the error
    return err
}
```

#### 🔹 Creating Custom Errors
Use `errors.New()` or `fmt.Errorf()`:
```go
import "errors"

err := errors.New("something went wrong")
```

Or with formatting:
```go
err := fmt.Errorf("failed to connect: %v", cause)
```

#### 🔹 Returning Errors from Functions
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

---

### **3️⃣ Code Examples**

#### Example 1: Basic Error Check
```go
package main

import (
    "errors"
    "fmt"
)

func greet(name string) (string, error) {
    if name == "" {
        return "", errors.New("name cannot be empty")
    }
    return "Hello, " + name, nil
}

func main() {
    msg, err := greet("")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println(msg)
}
```

#### Example 2: Use `fmt.Errorf` to Wrap Context
```go
err := fmt.Errorf("failed to send request: %w", err)
```
Use `%w` for error wrapping — handy when you want to check for specific error types using `errors.Is()` or `errors.As()`.

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Always check and handle errors – **ignoring them is a code smell**.
- ✅ Use `fmt.Errorf()` with `%w` to preserve error context.
- ✅ Create meaningful error messages that help debug the issue.
- ❌ Don’t panic unless it’s truly unrecoverable (like bad init configs).
- ❌ Avoid overusing `errors.New()` without context.

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a function `loadUser(id int) (string, error)` that:
- Returns `"User 1"` if `id == 1`
- Returns an error saying `"user not found"` for any other id

Then in `main()`, call this function with id `2`, and handle the error by printing:
```
Failed to load user: user not found
```

---

### **6️⃣ Further Reading & Resources**

- 📖 [Go Blog: Errors are Values](https://blog.golang.org/errors-are-values)
- 📖 [Effective Go – Error Handling](https://golang.org/doc/effective_go.html#errors)
- 📖 [pkg/errors for advanced patterns (archived)](https://github.com/pkg/errors)
- 📖 [Go 1.13 Error Wrapping Explained](https://go.dev/blog/go1.13-errors)
