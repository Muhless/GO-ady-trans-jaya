package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Golang")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a simple Go project example.")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	fmt.Println("Server berjalan di port 8080...")
	http.ListenAndServe(":8000", nil)
}
