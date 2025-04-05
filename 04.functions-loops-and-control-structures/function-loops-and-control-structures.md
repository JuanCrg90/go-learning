## **📌 Lesson 4: Functions, Loops, and Control Structures in Go**

### **1️⃣ Introduction**
Functions, loops, and control structures are core to writing logic in any language. Go’s syntax is clean and minimal, but powerful. It emphasizes readability and performance without sacrificing flexibility.

---

### **2️⃣ Core Concepts**

#### 🔹 Functions

Functions are declared using the `func` keyword:

```go
func greet(name string) string {
    return "Hello, " + name
}
```

- You can return multiple values:
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

- Functions can be anonymous or assigned to variables:
```go
add := func(a, b int) int {
    return a + b
}
```

---

#### 🔹 Loops

Go only has **one loop keyword**: `for`

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

- It also works like a `while` loop:
```go
i := 0
for i < 5 {
    fmt.Println(i)
    i++
}
```

- Infinite loop:
```go
for {
    fmt.Println("Looping forever...")
}
```

- Looping over slices or maps:
```go
numbers := []int{1, 2, 3}
for index, value := range numbers {
    fmt.Println(index, value)
}
```

---

#### 🔹 Control Structures

Go uses `if`, `else if`, and `else` blocks (no parentheses required!):

```go
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else {
    fmt.Println("Try again")
}
```

- You can declare variables inline with `if`:
```go
if age := 20; age >= 18 {
    fmt.Println("Adult")
}
```

- `switch` works like in C-style languages, but without needing `break`:
```go
day := "Tuesday"
switch day {
case "Monday":
    fmt.Println("Start of week")
case "Friday":
    fmt.Println("End of week")
default:
    fmt.Println("Midweek")
}
```

---

### **3️⃣  Common Pitfalls & Best Practices**

- ✅ **Use `range` for clean iteration over slices, maps, and arrays.**
- ✅ **Return multiple values for errors instead of exceptions.**
- ✅ **Prefer explicit, short functions – Go favors readability.**
- ❌ **Don’t rely on `break` in `switch`; it’s implicit.**
- ❌ **Avoid global variables for state – use function scope when possible.**

---

### ** 4️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:

1. Has a function `isEven(num int) bool` that returns `true` if the number is even.
2. Loops from 1 to 20 using a `for` loop.
3. Prints `"X is even"` or `"X is odd"` based on the result of `isEven`.

Example output:
```
1 is odd
2 is even
3 is odd
...
```

---

### **5️⃣ Further Reading & Resources**

- 📖 [Go Tour – Functions](https://go.dev/tour/moretypes/1)
- 📖 [Effective Go – Control Structures](https://golang.org/doc/effective_go.html#control-structures)
- 📖 [Go Playground (Try it online)](https://play.golang.org)
