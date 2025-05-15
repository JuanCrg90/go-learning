## **📌 Lesson 15: `sync` Package – WaitGroups and Mutexes**

### **1️⃣ Introduction**

Go provides the `sync` package for **low-level concurrency control**.

In this lesson, we’ll focus on:

* ✅ `sync.WaitGroup` – to wait for multiple goroutines to finish
* ⏳ `sync.Mutex` – to safely mutate shared data (coming later)

---

### **2️⃣ Core Concept: `sync.WaitGroup`**

#### 🔹 What is a WaitGroup?

A `WaitGroup` waits for a collection of goroutines to finish.

Think of it as:

* `.Add(1)` → register a new task
* `.Done()` → signal task completion
* `.Wait()` → block until all tasks complete

---

### **3️⃣ Example: Wait for 3 Workers**

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d started\n", id)
	// simulate work
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers done")
}
```

---

### 🔐 Why Use WaitGroup?

* Prevents using `time.Sleep()` as a crude wait tool
* Ensures main doesn't exit before goroutines complete
* Safer and more scalable than channel tricks for pure task coordination

---

### ✅ Best Practices

| Pattern                                         | Reason                                   |
| ----------------------------------------------- | ---------------------------------------- |
| Use `defer wg.Done()` at top of goroutine       | Ensures completion signal is always sent |
| Always pass `*sync.WaitGroup`                   | It must be passed by reference           |
| Call `wg.Add()` **before** launching goroutines | Avoid race conditions                    |

🔐 Golden Rule

🔁 For each goroutine you launch, call wg.Add(1) before starting it.
Inside that goroutine, always defer wg.Done() to signal completion.

---

### **4️⃣ Common Pitfalls**

- ❌ Forgetting `.Add()` or `.Done()` leads to hanging `.Wait()`
- ❌ Calling `.Add()` **inside** the goroutine (not safe if goroutine runs fast)
- ✅ Use `defer` to guarantee `.Done()` is called even if the function panics



---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:

1. Spawns 3 goroutines, each printing `"Routine X working"` and then `"Routine X done"`.
2. Uses a `sync.WaitGroup` to ensure the main function waits for all to finish.
3. Finally, prints `"Main done"` after all routines complete.

---

### **6️⃣ Further Reading**

- 📖 [Go by Example – WaitGroups](https://gobyexample.com/waitgroups)
- 📖 [`sync.WaitGroup` docs](https://pkg.go.dev/sync#WaitGroup)
