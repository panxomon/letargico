package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

// Hello World API
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3001" //localhost
	}

	helloHandler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!")
	}

	http.HandleFunc("/hello", helloHandler)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
