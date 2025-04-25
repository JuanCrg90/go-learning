## **📌 Lesson 6: Structs and Methods in Go**

### **1️⃣ Introduction**
Go doesn't have traditional classes, but it gives you **structs and methods** to model data and behavior. This is where Go’s minimalist but powerful object-oriented style comes in.

You’ll learn how to:
- Define your own data types with fields
- Attach behaviors to those types
- Compose types together (instead of using inheritance)

---

### **2️⃣ Core Concepts**

#### 🔹 Structs
Structs group fields together into a single compound type:
```go
type User struct {
    Name string
    Age  int
}
```

Create a struct instance:
```go
user := User{Name: "Juan", Age: 34}
```

Access fields:
```go
fmt.Println(user.Name)
```

#### 🔹 Structs are value types
Assigning a struct copies its value:
```go
u1 := User{Name: "Alice"}
u2 := u1
u2.Name = "Bob"
fmt.Println(u1.Name) // Alice (not Bob)
```

#### 🔹 Pointers to Structs
Use pointers if you want to mutate the original:
```go
func updateName(u *User, newName string) {
    u.Name = newName
}
```

#### 🔹 Methods on Structs
Attach methods to a type:
```go
func (u User) Greet() string {
    return "Hello, " + u.Name
}
```

If you need to modify the receiver:
```go
func (u *User) CelebrateBirthday() {
    u.Age++
}
```

---

### **3️⃣ Code Examples**

#### Example 1: Struct + Method
```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

#### Example 2: Mutating with Pointer Receiver
```go
type Counter struct {
    Value int
}

func (c *Counter) Increment() {
    c.Value++
}
```

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Use **value receivers** if your method doesn't modify the receiver.
- ✅ Use **pointer receivers** when you need to modify the struct or to avoid copying.
- ✅ **Favor composition over inheritance** – Go doesn't have inheritance by design.
- ❌ Don't mix pointer and value receivers on the same struct unless necessary — it leads to confusion.
- ❌ Avoid exporting struct fields unless they need to be accessed outside the package.

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that defines a `Book` struct with the following fields:
- `Title` (string)
- `Author` (string)
- `Pages` (int)

Create two methods:
- `Summary() string` – returns a string like `"Title by Author, 123 pages"`
- `AddPages(extra int)` – adds pages to the book (use a pointer receiver)

Then:
- Create a book, print its summary, add 50 pages, and print the new summary.

---

### **6️⃣ Further Reading & Resources**

📖 [Go by Example – Structs](https://gobyexample.com/structs)
📖 [Effective Go – Structs and Methods](https://golang.org/doc/effective_go.html#structs)
📖 [Go Tour – Methods](https://go.dev/tour/methods/1)
