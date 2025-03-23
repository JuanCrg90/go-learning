## **📌 Lesson 3: Understanding Variables, Constants, and Data Types in Go**

### **1️⃣ Introduction**
In Go, variables and constants are fundamental building blocks. Go is a statically-typed language, meaning every variable has a specific type at compile-time. Understanding how Go handles data types will help you write efficient, bug-free code.

---

### **2️⃣ Core Concepts**
#### 🔹 Declaring Variables
Go provides multiple ways to declare variables:
```go
// Explicit type declaration
var name string = "Juan"

// Implicit type inference
age := 34  // Automatically inferred as int

// Multiple variable declaration
var x, y int = 10, 20

// Zero values (default values for uninitialized variables)
var uninitialized string // Default is an empty string ("")
var number int           // Default is 0
```

#### 🔹 Constants
Constants are immutable values declared using the `const` keyword.
```go
const PI float64 = 3.14159
const Greeting = "Hello, Go!"
```
- Constants cannot be declared using `:=` (short assignment).
- They are evaluated at compile time.

#### 🔹 Data Types
Go has several built-in types:

| **Category**  | **Types**  | **Example**  |
|--------------|-----------|-------------|
| **Boolean**  | `bool`    | `true`, `false`  |
| **Integer**  | `int`, `int8`, `int16`, `int32`, `int64`  | `var num int = 42` |
| **Unsigned Integer**  | `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `var u uint = 255` |
| **Float**  | `float32`, `float64` | `var f float64 = 3.14` |
| **Complex**  | `complex64`, `complex128` | `var c complex128 = 1 + 2i` |
| **String**  | `string` | `var s string = "GoLang"` |
| **Array**  | `[size]Type` | `var arr [3]int = [3]int{1, 2, 3}` |
| **Slice**  | `[]Type` | `var sl []int = []int{1, 2, 3}` |
| **Map**  | `map[KeyType]ValueType` | `var m map[string]int = map[string]int{"a": 1, "b": 2}` |
| **Struct**  | `struct{}` | `type Person struct { Name string; Age int }` |

#### 🔹 Type Conversion
Go does **not** allow implicit type conversion. You must explicitly convert types:
```go
var i int = 10
var f float64 = float64(i) // Convert int to float64
var s string = strconv.Itoa(i) // Convert int to string
```

---

### **3️⃣ Code Examples**
#### Example 1: Working with Variables
```go
package main

import "fmt"

func main() {
    var name string = "Juan"
    age := 34
    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
```

#### Example 2: Constants and Type Inference
```go
package main

import "fmt"

const Pi = 3.14159

func main() {
    const Greeting = "Welcome to Golang!"
    fmt.Println(Greeting, "Pi =", Pi)
}
```

#### Example 3: Type Conversion
```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    var num int = 42
    str := strconv.Itoa(num) // Convert int to string
    fmt.Println("String:", str)
}
```

---

### **4️⃣ Common Pitfalls & Best Practices**
- ✅ **Use short declaration (`:=`) for local variables**, but `var` for package-level variables.
- ✅ **Prefer `const` over `var` when the value won’t change.**
- ✅ **Be explicit with type conversions** (Go doesn’t allow implicit conversions).
- ❌ **Don’t use uninitialized variables unnecessarily**, since Go provides default values.
- ❌ **Avoid excessive type conversions**, as they can impact performance.

---

### **5️⃣ Hands-on Exercise**
👉 **Challenge**: Write a Go program that:
- Declares a `const` named `BaseSalary` with a value of `50000`.
- Declares a `var` named `bonus` initialized to `5000`.
- Calculates the `totalSalary` by adding `BaseSalary + bonus`.
- Converts the `totalSalary` to a string and prints a message like:
  `"The total salary is 55000"`

**Try it out! Once you complete it, share your code and I’ll review it. 🚀**

---

### **6️⃣ Further Reading & Resources**
- 📖 [Go Tour - Variables](https://tour.golang.org/basics/8)
- 📖 [Go Docs - Constants](https://golang.org/ref/spec#Constants)
- 📖 [Effective Go - Declarations](https://golang.org/doc/effective_go.html#declarations)
