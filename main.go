package main

import 
(
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
