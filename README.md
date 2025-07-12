# Go Website Status Checker

This is a simple program written in Go. It checks if a website is working by sending a request and reading the response. In this example, it checks the Google website.

## What It Does
- Sends a request to `https://google.com`
- If the website is up, it prints a message and shows the response status code
- If it cannot reach the website, it prints an error message

## Code Explanation
```go
package main
import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Could not reach the website.")
		return
	}

	fmt.Println("Website is up!")
	fmt.Println("Status code:", resp.StatusCode)
}
```

## How to Run
1. Make sure Go is installed on your computer
2. Save the code in a file named `main.go`
3. Open a terminal or command prompt and run:
```bash
go run main.go
```
4. You will see either:
   - "Website is up!" and the status code
   - Or "Could not reach the website." if something goes wrong

## Requirements
- Go programming language installed
- Internet connection

## Why This Project
This is a good starting point for beginners who want to:
- Learn Go basics
- Work with HTTP requests
- Practice writing simple and useful programs

## Author
Chandana Alagod

