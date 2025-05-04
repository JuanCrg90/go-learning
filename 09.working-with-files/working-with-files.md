## **📌 Lesson 9: Working with Files in Go (`os`, `io`, `bufio`)**

### **1️⃣ Introduction**
File handling is essential in real-world applications: from reading configuration files to storing logs, user data, or cached content.

Go provides powerful and straightforward tools via the `os`, `io`, and `bufio` packages for:
- Reading/writing files
- Creating and deleting files
- Scanning line-by-line content

---

### **2️⃣ Core Concepts**

---

#### 🔹 Opening Files
```go
file, err := os.Open("example.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()
```

Use `defer` to ensure the file is closed after reading.

---

#### 🔹 Reading Files (Line by Line)
```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
```

Check for scan errors:
```go
if err := scanner.Err(); err != nil {
    log.Fatal(err)
}
```

---

#### 🔹 Reading Entire File as Bytes
```go
content, err := os.ReadFile("example.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(content))
```

---

#### 🔹 Writing to a File
```go
file, err := os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

_, err = file.WriteString("Hello, file!\n")
```

Append mode:
```go
file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
```

---

#### 🔹 Removing Files
```go
err := os.Remove("output.txt")
if err != nil {
    log.Fatal(err)
}
```

---

### **3️⃣ Code Examples**

#### Example 1: Write and Read File
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    err := os.WriteFile("hello.txt", []byte("Hello, Juan Carlos!"), 0644)
    if err != nil {
        panic(err)
    }

    data, err := os.ReadFile("hello.txt")
    if err != nil {
        panic(err)
    }

    fmt.Println(string(data))
}
```

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Always check for `err` when working with files.
- ✅ Use `defer file.Close()` immediately after `Open/Create`.
- ✅ Use `bufio` for large files or line-by-line scanning.
- ❌ Don’t hardcode paths – use `filepath` for cross-platform compatibility.
- ❌ Avoid loading huge files with `ReadFile` unless memory is not a concern.

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:
1. Creates a file called `log.txt`.
2. Writes the following lines to it:
```
Log started.
User: Juan Carlos
Status: Active
```
3. Reads the file line by line and prints each line to the console.

**Tip**: Use `bufio.NewScanner` to read line by line.

---

### **6️⃣ Further Reading & Resources**

- 📖 [Go by Example – File I/O](https://gobyexample.com/reading-files)
- 📖 [os package docs](https://pkg.go.dev/os)
- 📖 [bufio package docs](https://pkg.go.dev/bufio)
- 📖 [io/ioutil deprecated notice](https://pkg.go.dev/io/ioutil)
