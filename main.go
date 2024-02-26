// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, Sock Shop!")

	// Your main application logic goes here
	// For example, start a web server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Sock Shop!")
	})

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
