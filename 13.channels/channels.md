## **📌 Lesson 13: Channels – Communicating Between Goroutines**

### **1️⃣ Introduction**

Goroutines give us concurrency. **Channels** give us coordination.

> A **channel** is a typed conduit through which you can send and receive values **between goroutines** safely.

---

### **2️⃣ Core Concepts**

---

#### 🔹 Creating a Channel

```go
ch := make(chan string)
```

You can only send/receive the declared type (`string` in this case).

---

#### 🔹 Sending and Receiving

```go
go func() {
    ch <- "hello" // send
}()

msg := <-ch // receive
fmt.Println(msg)
```

✅ Both send and receive **block** until the other side is ready — this is how Go avoids race conditions without locks.

---

#### 🔹 Directional Channels

You can declare **send-only** or **receive-only** channels:

```go
func sendMsg(ch chan<- string) {
    ch <- "data"
}

func receiveMsg(ch <-chan string) {
    msg := <-ch
    fmt.Println(msg)
}
```

---

#### 🔹 Buffered Channels

If you want to **decouple** send/receive timing, use a buffer:

```go
ch := make(chan string, 2)
ch <- "A"
ch <- "B"
// buffer has 2 items, no blocking yet
```

---

#### 🔹 Closing Channels

```go
close(ch)
```

You can loop over a channel with `range` until it's closed:

```go
for msg := range ch {
    fmt.Println(msg)
}
```

---

### **3️⃣ Code Example: Parallel Tasks with Channel**

```go
package main

import (
    "fmt"
)

func worker(id int, ch chan string) {
    ch <- fmt.Sprintf("Worker %d done", id)
}

func main() {
    ch := make(chan string)

    for i := 1; i <= 3; i++ {
        go worker(i, ch)
    }

    for i := 0; i < 3; i++ {
        msg := <-ch
        fmt.Println(msg)
    }

    fmt.Println("All workers done")
}
```

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Prefer channels for coordination over shared memory
- ✅ Use **buffered channels** only when you need decoupled communication
- ✅ Always **close channels** when done sending (especially in producers)
- ❌ Never read from a **nil** or closed channel
- ❌ Avoid deadlocks: every send must have a matching receive

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:

1. Spawns 3 goroutines (workers).
2. Each goroutine sends a message like `"Routine X done"` to a shared channel.
3. The main function reads all 3 messages and prints them.

**Expected output (order may vary):**

```
Routine 1 done
Routine 3 done
Routine 2 done
Main function done
```

✅ No `time.Sleep` — only use channels.

---

### **6️⃣ Further Reading & Resources**

- 📖 [Go Tour – Channels](https://go.dev/tour/concurrency/2)
- 📖 [Effective Go – Channels](https://golang.org/doc/effective_go.html#channels)
- 📖 [Go Blog – Share Memory by Communicating](https://go.dev/blog/share-memory-by-communicating)
