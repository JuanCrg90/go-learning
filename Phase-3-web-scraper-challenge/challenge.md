## 🚀 Project: Concurrent Web Scraper

### **🎯 Objective**

Build a CLI Go program that:

* Fetches content from multiple URLs concurrently
* Measures response times
* Prints results in a clean, readable format

---

## ✅ Suggested Features

| Feature               | Description                                  |
| --------------------- | -------------------------------------------- |
| 🧵 Goroutines         | Launch a goroutine per URL                   |
| 🔗 Channels           | Collect results (URL, status code, duration) |
| ⌛ WaitGroup           | Wait for all requests to complete            |
| ⏱️ Context (optional) | Add timeout for slow URLs                    |
| 📊 Output             | Show status and time per site in console     |

---

## 📦 Suggested Project Structure

```
web-scraper/
├── main.go
├── fetcher.go   # logic to fetch URLs
├── types.go     # result structs (optional)
```

---

## 🧪 Starter URLs

Use a list like:

```go
urls := []string{
	"https://golang.org",
	"https://httpbin.org/delay/2",
	"https://example.com",
}
```

---

## ✨ Bonus Ideas (Optional Enhancements)

* Retry failed requests
* Save output as JSON or CSV
* Use a context to enforce global timeout
* Color-code fast/slow responses in output
