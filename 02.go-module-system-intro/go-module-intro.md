# **Lesson: Go's Module System (`go mod init`, `go get`)**

## **1. Introduction**

Go modules are Go’s dependency management system, introduced in Go 1.11 and made the default in Go 1.14. They provide a way to manage dependencies, version control, and package organization without relying on the old `GOPATH` workspace.

### **Why are Go modules important?**
- **Reproducibility**: Modules allow projects to have a fixed set of dependencies, ensuring consistent builds.
- **Versioning**: They enable versioning of dependencies using semantic versioning (`v1.2.3` format).
- **Ease of Use**: Developers no longer need to keep all Go projects inside `GOPATH`, making Go development more flexible.

**Question for you:**
Have you worked with dependency management tools before (e.g., `npm` for JavaScript, `Bundler` for Ruby, or `Cargo` for Rust)? If so, how do you think they compare to Go’s approach?

---

## **2. Core Concepts**

Here are the key components of Go modules:

### **1. `go.mod` (Module Definition File)**
- This file is created when you initialize a Go module using `go mod init <module-name>`.
- It records the module name and Go version.
- It tracks dependencies explicitly added via `go get`.

**Example `go.mod` file:**
```go
module github.com/yourusername/myproject

go 1.21
```

### **2. `go.sum` (Dependency Checksum File)**
- Stores cryptographic checksums for each dependency to ensure integrity.
- Helps prevent tampering by verifying that downloaded dependencies match expected versions.

**Example `go.sum` file:**
```
github.com/gorilla/mux v1.8.0 h1:abcxyz...
github.com/gorilla/mux v1.8.0/go.mod h1:123456...
```

### **3. `go mod init <module-name>` (Initialize a New Module)**
- Creates a new `go.mod` file in the current directory.
- The module name should ideally be a URL-like identifier (e.g., `github.com/user/project`).

### **4. `go get <package>` (Adding Dependencies)**
- Fetches a dependency and updates `go.mod` and `go.sum`.
- Example: `go get github.com/gorilla/mux@latest`

### **5. `go mod tidy` (Clean Up Unused Dependencies)**
- Removes unused dependencies from `go.mod` and `go.sum`.
- Ensures that only the necessary dependencies remain in the project.

---

## **3. Step-by-Step Guide**

### **Step 1: Initialize a New Go Module**
1. Open a terminal and navigate to your project directory.
2. Run:
   ```sh
   go mod init github.com/yourusername/myproject
   ```
3. This creates a `go.mod` file.

### **Step 2: Add Dependencies**
1. Install a package, such as `gorilla/mux` for routing:
   ```sh
   go get github.com/gorilla/mux@latest
   ```
2. This updates `go.mod` and `go.sum`.

### **Step 3: Use the Dependency in Your Code**
Create a file `main.go`:
```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Go Modules!")
	})

	http.ListenAndServe(":8080", r)
}
```

### **Step 4: Clean Up Unused Dependencies**
Run:
```sh
go mod tidy
```
This removes unnecessary dependencies.

---

## **4. Common Pitfalls & Best Practices**

### **Common Mistakes**
❌ **Forgetting to initialize the module** → Always start with `go mod init`.
❌ **Using `go get` without a version** → Always use `@latest` or specify a version for stability.
❌ **Deleting `go.sum`** → This file ensures integrity; don’t delete it manually.
❌ **Committing unnecessary dependencies** → Run `go mod tidy` regularly.

### **Best Practices**
✅ **Use semantic versioning** when installing dependencies (`@v1.2.3`).
✅ **Run `go mod tidy` before committing** to avoid unused dependencies.
✅ **Use `replace` directives cautiously** (useful for local module development).
✅ **Keep dependencies minimal**—only install what you need.

---

## **5. Hands-on Exercise**

**Objective:**
1. Initialize a Go module.
2. Add a dependency (`gorilla/mux`).
3. Write a simple HTTP server using the dependency.
4. Run `go mod tidy` and observe changes.

### **Instructions:**
1. Create a new directory and navigate into it:
   ```sh
   mkdir go-modules-demo && cd go-modules-demo
   ```
2. Initialize a new Go module:
   ```sh
   go mod init github.com/yourusername/go-modules-demo
   ```
3. Install `gorilla/mux`:
   ```sh
   go get github.com/gorilla/mux@latest
   ```
4. Create `main.go`:
   ```go
   package main

   import (
   	"fmt"
   	"net/http"

   	"github.com/gorilla/mux"
   )

   func main() {
   	r := mux.NewRouter()
   	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   		fmt.Fprintln(w, "Hello, Go Modules!")
   	})

   	http.ListenAndServe(":8080", r)
   }
   ```
5. Run the server:
   ```sh
   go run main.go
   ```
6. Open `http://localhost:8080` in your browser.
7. Run `go mod tidy` and check if `go.mod` or `go.sum` changed.

Let me know once you finish the exercise so we can review it together!

---

## **6. Further Reading & Resources**

- [Go Modules Reference](https://go.dev/ref/mod)
- [Go Blog: Using Go Modules](https://go.dev/blog/using-go-modules)
- **Book:** *The Go Programming Language* by Alan Donovan & Brian Kernighan (Ch. 10 covers modules)


### **Why `go mod tidy` Didn't Change Anything?**
1. **All dependencies are in use**: `go mod tidy` removes unused dependencies, but since `gorilla/mux` is actually imported and used in your `main.go`, it remains in `go.mod` and `go.sum`.
2. **No extraneous dependencies**: You didn’t install extra packages that weren't used in your code.
3. **No manual modifications**: If you had manually removed an import or edited `go.mod`, `go mod tidy` might have updated the files.

### **When Would `go mod tidy` Change Things?**
- If you had installed a dependency but never used it (`go mod tidy` would remove it).
- If you had unused indirect dependencies (`go mod tidy` would clean them up).
- If you removed an import statement but didn't remove the dependency manually.

**Want to test this?**
Try running:
```sh
go get github.com/sirupsen/logrus
```
Then check `go.mod`—you'll see `logrus` added. Now run:
```sh
go mod tidy
```
Since you never imported `logrus` in your code, `go mod tidy` will remove it.
