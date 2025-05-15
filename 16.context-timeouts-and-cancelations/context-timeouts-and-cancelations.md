
## **📌 Lesson 16: `context` Package – Timeouts & Cancellations**

### **1️⃣ Introduction**

The `context` package is designed to control the **lifecycle of operations** across goroutines. It allows you to:

* Cancel goroutines when no longer needed
* Set deadlines and timeouts
* Propagate cancellation across API boundaries

---

### **2️⃣ Core Concepts**

#### 🔹 What is a Context?

A `context.Context` carries deadlines, cancellation signals, and other request-scoped values.

---

#### 🔹 Creating a Cancellable Context

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()
```

You pass `ctx` to functions or goroutines. When you call `cancel()`, anything listening to that `ctx` stops.

---

#### 🔹 Timeout or Deadline

```go
ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
defer cancel()
```

Or use a fixed time:

```go
deadline := time.Now().Add(1 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
```

---

#### 🔹 Listening for Cancellation

Inside a goroutine:

```go
select {
case <-ctx.Done():
    fmt.Println("Cancelled:", ctx.Err())
}
```

---

### **3️⃣ Example: Cancelling a Worker**

```go
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
```

```go
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("Main done")
}
```

---

### ✅ Best Practices

| Rule                                                           | Why                    |
| -------------------------------------------------------------- | ---------------------- |
| Always `defer cancel()`                                        | Avoid resource leaks   |
| Pass `context.Context` as first parameter in functions         | Go convention          |
| Check `ctx.Done()` in loops or blocking code                   | Ensure graceful exit   |
| Use `context.WithTimeout` or `WithDeadline` for external calls | Prevent stuck requests |

---

### **4️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:

1. Starts a goroutine that prints “Routine working...” every 200ms.
2. Cancels that routine automatically after 1 second using `context.WithTimeout`.
3. Prints “Routine cancelled” once it stops, and “Main done” at the end.

---

### **5️⃣ Further Reading**

- 📖 [`context` package docs](https://pkg.go.dev/context)
- 📖 [Go Blog – Context](https://blog.golang.org/context)
