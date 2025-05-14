## **📌 Lesson 14: `select` Statement – Handling Multiple Channels**

### **1️⃣ Introduction**

Go’s `select` statement lets a goroutine **wait on multiple channel operations** at once — it’s like a `switch` but for channels.

This is incredibly useful when:

* You want to wait on **whichever channel responds first**
* You're handling **timeouts**, **fan-in/fan-out**, or **cancellation**

---

### **2️⃣ Core Concepts**

---

#### 🔹 Basic `select` Syntax

```go
select {
case msg1 := <-ch1:
    fmt.Println("Received from ch1:", msg1)
case msg2 := <-ch2:
    fmt.Println("Received from ch2:", msg2)
}
```

* Whichever channel is ready **first** will proceed.
* If **multiple are ready**, one is chosen **at random**.
* If **none are ready**, `select` blocks **until one becomes ready**.

---

#### 🔹 `default` Case (Non-blocking Select)

```go
select {
case msg := <-ch:
    fmt.Println("Received:", msg)
default:
    fmt.Println("No message received")
}
```

✅ Useful to avoid blocking or implement polling.

---

### **3️⃣ Example: Competing Channels**

```go
package main

import (
    "fmt"
    "time"
)

func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)

    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "Message from ch1"
    }()

    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "Message from ch2"
    }()

    select {
    case msg := <-ch1:
        fmt.Println("Received:", msg)
    case msg := <-ch2:
        fmt.Println("Received:", msg)
    }
}
```

> Output: `"Message from ch1"` — because it arrives first.

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Use `select` to wait for **any of multiple results**
- ✅ Use a `default` case **sparingly** – it makes `select` non-blocking
- ✅ Use `select` in infinite `for` loops for **event-driven goroutines**
- ❌ Don’t read from nil channels — it causes a **permanent block**

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:

1. Creates two goroutines, each sending to a separate channel after a random delay.
2. Uses a `select` to print whichever message arrives first.
3. Print `"Main done"` after receiving one message.

🧠 Hint: Use `time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)` for random delay (you’ll need to `import "math/rand"`).

---

### **6️⃣ Further Reading**

- 📖 [Go Tour – `select`](https://go.dev/tour/concurrency/5)
- 📖 [Go Blog – Timeout Patterns](https://blog.golang.org/concurrency-patterns-timing-out-and)
