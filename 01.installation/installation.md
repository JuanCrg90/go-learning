# Lesson 1: Installing Go and Setting Up Your Workspace

---

## 1. Introduction

Golang (or simply Go) is a statically typed, compiled programming language designed for efficiency, concurrency, and ease of use. Properly installing Go and setting up your development environment is crucial because:

- It ensures that you can run and compile Go programs without issues.
- A well-configured workspace allows for better package management and dependency handling.
- Go’s tooling (e.g., `go mod`, `go build`, `go run`) depends on a correct setup to function smoothly.

By the end of this lesson, you’ll have Go installed, your workspace configured, and a simple Go program running on your machine.

---

## 2. Core Concepts

Before we begin the installation process, let’s clarify key concepts:

### ✅ **Go Installation Directory**
- When you install Go, it places the Go compiler (`go`), the standard library, and other tools in a system-wide directory.
- On Linux/macOS, it's usually `/usr/local/go`.
- On Windows, it's typically `C:\Go`.

### ✅ **GOPATH vs GOROOT**
- `GOROOT`: This is where Go itself is installed (system-wide). You should **not** modify this.
- `GOPATH`: This is where your Go workspace (projects, source code, and dependencies) lives. You can set it manually, but since Go 1.11+, **modules are preferred over GOPATH**.

### ✅ **Go Modules (Preferred for New Projects)**
- Go modules (`go mod init`, `go mod tidy`) replace the traditional GOPATH structure, making dependency management easier.
- You don’t need to place projects inside a specific folder (`GOPATH/src`) anymore.

---

## 3. Step-by-Step Guide

Now, let’s install Go based on your operating system.

### 🔹 **Windows Installation**

1. Download the official installer from [Go’s website](https://go.dev/dl/).
2. Run the installer and follow the on-screen instructions (default location: `C:\Go`).
3. Add Go’s binary directory (`C:\Go\bin`) to your `PATH` environment variable (if not set automatically).
4. Open Command Prompt (`cmd`) or PowerShell and verify installation:
   ```powershell
   go version
   ```

### 🔹 **macOS Installation**

#### **Method 1: Using Homebrew (Recommended)**
1. Install Go via Homebrew:
   ```sh
   brew install go
   ```
2. Verify installation:
   ```sh
   go version
   ```

#### **Method 2: Manual Installation**
1. Download the macOS `.pkg` installer from [Go’s official website](https://go.dev/dl/).
2. Run the installer and follow the instructions.
3. Check installation:
   ```sh
   go version
   ```

### 🔹 **Linux Installation**

1. Download the latest Go tarball:
   ```sh
   wget https://go.dev/dl/go1.21.4.linux-amd64.tar.gz
   ```
2. Extract it to `/usr/local/`:
   ```sh
   sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
   ```
3. Add Go to your `PATH`:
   ```sh
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```
4. Verify installation:
   ```sh
   go version
   ```

---

## 4. Common Pitfalls & Best Practices

🔴 **Common Mistakes**
- Not adding Go’s binary directory to `PATH`, causing `go` commands to not be recognized.
- Confusing `GOROOT` with `GOPATH` and manually modifying `GOROOT` (which you should not).
- Not initializing a Go module (`go mod init`) before writing code, leading to dependency issues.

✅ **Best Practices**
- Prefer **Go modules** (`go mod init`) over `GOPATH`.
- Use an IDE like **VS Code** with the Go plugin for better development experience.
- Always verify installation with `go version` before writing code.

---

## 5. Hands-on Exercise

### 🏆 **Exercise: Install Go, Set Up Workspace, and Run Your First Program**

#### **Step 1: Verify Installation**
Run the following command to confirm Go is installed:
```sh
go version
```
If it prints a version (e.g., `go version go1.21.4`), you’re good to go!

#### **Step 2: Create Your First Go Project**
1. Create a new directory for your project:
   ```sh
   mkdir myfirstgo && cd myfirstgo
   ```
2. Initialize a Go module:
   ```sh
   go mod init myfirstgo
   ```
3. Create a file named `main.go` and open it in an editor.

#### **Step 3: Write a Simple Go Program**
Copy and paste this into `main.go`:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Golang!")
}
```

#### **Step 4: Run the Program**
Run your program with:
```sh
go run main.go
```
Expected output:
```
Hello, Golang!
```

#### **Step 5: Compile and Run**
To build an executable:
```sh
go build -o myapp
```
Now, run the compiled binary:
```sh
./myapp   # (Linux/macOS)
myapp.exe # (Windows)
```

---

## 6. Further Reading & Resources

📖 **Official Go Documentation:**
- [Go Installation Guide](https://go.dev/doc/install)
- [Go Modules Overview](https://go.dev/doc/modules)
- [Effective Go](https://go.dev/doc/effective_go)

📺 **Video Tutorials:**
- [Golang Crash Course – FreeCodeCamp](https://www.youtube.com/watch?v=YS4e4q9oBaU)

📚 **Books:**
- *The Go Programming Language* – Alan A. A. Donovan & Brian W. Kernighan
- *Go in Action* – William Kennedy
