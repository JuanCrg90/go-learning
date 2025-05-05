## **📌 Lesson 11: Writing Unit Tests with the `testing` Package**

### **1️⃣ Introduction**

Go has a built-in testing framework in the `testing` package. Unlike other languages that rely on external libraries, **Go treats tests as first-class citizens** using files with the `_test.go` suffix.

With `go test`, you can:

* Write and run unit tests
* Benchmark performance
* Create table-driven tests
* Catch regressions and bugs early

---

### **2️⃣ Core Concepts**

---

#### 🔹 A Basic Test

```go
// file: math_test.go
package math

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5

    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

* Tests must start with `Test` and take `*testing.T` as a parameter.
* Use `t.Errorf()` or `t.Fatal()` to report failures.

---

#### 🔹 Run the Tests

```sh
go test
```

Optional flags:

* `-v` (verbose)
* `-run=TestName` to run a specific test

---

#### 🔹 Table-Driven Tests (Go idiom)

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        a, b int
        want int
    }{
        {1, 1, 2},
        {2, 3, 5},
        {0, 0, 0},
    }

    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, got, tt.want)
        }
    }
}
```

---

#### 🔹 Organizing Tests

* Tests live in the **same package** or in a `_test` suffix package for black-box testing.
* Files must be named `*_test.go`.

---

#### 🔹 Additional Tools

* `t.Helper()` marks helper functions so line numbers in errors point to the actual test.
* `t.Run()` lets you nest subtests.
* Use `go test -cover` to check test coverage.

---

### **3️⃣ Code Examples**

#### Example: Testing a Book Summary Function

```go
// book.go
package book

type Book struct {
    Title  string
    Author string
}

func (b Book) Summary() string {
    return b.Title + " by " + b.Author
}
```

```go
// book_test.go
package book

import "testing"

func TestSummary(t *testing.T) {
    b := Book{"The Hobbit", "J.R.R. Tolkien"}
    want := "The Hobbit by J.R.R. Tolkien"
    got := b.Summary()

    if got != want {
        t.Errorf("Summary() = %q; want %q", got, want)
    }
}
```

Run with:

```sh
go test ./...
```

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Always test public methods and exported behavior.
- ✅ Use **table-driven tests** for scalability.
- ✅ Keep each test **independent and deterministic**.
- ❌ Don’t test unexported functions directly from other packages.
- ❌ Avoid relying on external state (files, DBs) in unit tests without isolation.

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a small testable Go program in two files:

**`calculator.go`**

* Implement a `Sum(a, b int) int` function.

**`calculator_test.go`**

* Write a table-driven test to verify the `Sum()` function works for:

  * `1 + 1 = 2`
  * `5 + 7 = 12`
  * `-3 + 3 = 0`

Run it with `go test`.

---

### **6️⃣ Further Reading & Resources**

- 📖 [Go by Example – Testing](https://gobyexample.com/testing)
- 📖 [`testing` package docs](https://pkg.go.dev/testing)
