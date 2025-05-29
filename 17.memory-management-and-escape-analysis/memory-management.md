## **📌 Lesson 17: Memory Management & Escape Analysis in Go**

### **1️⃣ Introduction**

Go is a garbage-collected language, but understanding **where** your variables live — **stack** vs **heap** — is critical for writing **high-performance** code.

This is where **escape analysis** comes in:

> It tells the Go compiler whether a variable can be safely kept on the stack or if it “escapes” to the heap.

---

### **2️⃣ Stack vs Heap Recap**

| Location               | Characteristics                                            |
| ---------------------- | ---------------------------------------------------------- |
| **Stack**              | Fast allocation, limited scope (function-local)            |
| **Heap**               | Slower, used when variables need to outlive function calls |
| **Garbage collected?** | Stack ❌ No, Heap ✅ Yes                                     |


---

### **3️⃣ When Does a Variable Escape?**

A variable escapes to the heap when:

* It's returned by a function
* It’s captured by a goroutine or closure
* It's referenced by a pointer used outside the function’s scope

**🧠 Bonus Tip: Keep This Rule of Thumb**
| Return Type                        | Likely Heap Allocation? |
| ---------------------------------- | ----------------------- |
| Return by **value**                | ❌ No                    |
| Return **pointer to local**        | ✅ Yes                   |
| Use in a **closure**               | ✅ Yes                   |
| Store in a **struct that escapes** | ✅ Yes                   |

---

#### 🔹 Example: Stack Allocation

```go
func foo() int {
    x := 42   // stays on the stack
    return x
}
```

---

#### 🔹 Example: Escape to Heap

```go
func bar() *int {
    x := 42
    return &x // escapes to heap
}
```

Because `x`'s memory is needed **after the function ends**, it must be on the heap.

---

### **4️⃣ How to Analyze It**

Use the `-gcflags` escape diagnostic:

```bash
go build -gcflags="-m"
```

Or to see it during run:

```bash
go run -gcflags="-m" main.go
```

It will output lines like:

```
./main.go:10:6: moved to heap: x
```

---

### **5️⃣ Why It Matters**

* Stack allocations are **faster and cheaper**
* Heap allocations cause **GC pressure** and **slow down your app**
* Escape analysis helps you **write tighter, high-performance code** (especially for tight loops or high-volume services)

---

### **6️⃣ Hands-on Exercise**

👉 **Challenge**: Write two small functions:

1. One that returns an `int` (value)
2. One that returns a pointer to an `int`

Then run:

```bash
go run -gcflags="-m" main.go
```

And observe which values escape to the heap.

---

### **7️⃣ Further Reading**

- 📖 [Go Blog – Go Memory Management](https://blog.golang.org/go-memstats)
- 📖 [Uber Go Style Guide – Performance](https://github.com/uber-go/guide/blob/master/style.md#performance)
