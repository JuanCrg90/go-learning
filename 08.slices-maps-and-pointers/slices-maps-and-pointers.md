## **📌 Lesson 8: Working with Slices, Maps, and Pointers in Go**

### **1️⃣ Introduction**
In Go, **slices** and **maps** are the workhorses for managing collections of data, while **pointers** give you fine-grained control over memory and performance.
Together, they unlock efficient, powerful, and clean code.

---

### **2️⃣ Core Concepts**

---

#### 🔹 Slices

- A **slice** is a dynamically-sized, flexible view into the elements of an array.
- Unlike arrays, slices **don’t have a fixed length**.

Creating a slice:
```go
numbers := []int{1, 2, 3}
```

Key slice operations:
```go
fmt.Println(numbers[0])     // Access first element
fmt.Println(len(numbers))   // Length
numbers = append(numbers, 4, 5) // Add elements
```

Slices from arrays:
```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2 3 4]
```

---

#### 🔹 Maps

- A **map** is a collection of key-value pairs.

Creating a map:
```go
m := map[string]int{
    "a": 1,
    "b": 2,
}
```

Basic map operations:
```go
value := m["a"]
m["c"] = 3           // Add or update
delete(m, "b")       // Remove
_, ok := m["d"]      // Check existence
```

---

#### 🔹 Pointers

- A **pointer** holds the memory address of a value.
- Use `*` to dereference (read/write the value) and `&` to take the address.

Example:
```go
x := 5
p := &x    // p is a pointer to x
fmt.Println(*p) // dereference pointer -> 5
*p = 10
fmt.Println(x)  // x is now 10
```

Pointers to slices and maps:
- Slices and maps are **reference types** — passing them around already shares the underlying data.
- You rarely need pointers to slices or maps unless you want to reassign them.

---

### **3️⃣ Code Examples**

#### Example 1: Basic Slice
```go
fruits := []string{"apple", "banana", "cherry"}
for _, fruit := range fruits {
    fmt.Println(fruit)
}
```

#### Example 2: Basic Map
```go
ages := map[string]int{
    "Alice": 25,
    "Bob": 30,
}

for name, age := range ages {
    fmt.Println(name, "is", age, "years old")
}
```

#### Example 3: Pointer Function
```go
func increment(x *int) {
    *x++
}

func main() {
    num := 5
    increment(&num)
    fmt.Println(num) // 6
}
```

---

### **4️⃣ Common Pitfalls & Best Practices**

- ✅ Use `append()` to grow slices — don’t manually resize arrays.
- ✅ Always check existence when reading from maps (`_, ok := m[key]`).
- ✅ Use pointers for **shared, mutable** state (but not unnecessarily).
- ❌ Be cautious with slices backed by arrays — changes can affect both.
- ❌ Avoid overusing pointers where value semantics would be clearer.

---

### **5️⃣ Hands-on Exercise**

👉 **Challenge**: Write a program that:
- Creates a `map[string]int` representing a simple inventory (`"apple": 5, "banana": 2`).
- Write a function `updateInventory` that receives the inventory map and an item name, and **increments** the item count by 1.
- If the item doesn’t exist, **add it** with a count of 1.
- Test updating `"apple"`, and adding a new `"orange"`.

Example output:
```
map[apple:6 banana:2 orange:1]
```

---

### **6️⃣ Further Reading & Resources**

- 📖 [Go by Example – Slices](https://gobyexample.com/slices)
- 📖 [Go by Example – Maps](https://gobyexample.com/maps)
- 📖 [Go Tour – Pointers](https://go.dev/tour/moretypes/1)
