package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// helloHandler handles the root path
func helloHandler(w http.ResponseWriter, r *http.Request) {
	greeting := os.Getenv("GREETING")
	fmt.Fprintln(w, greeting)
}

// notFoundHandler handles 404 errors
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 - Not Found", http.StatusNotFound)
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define your routes
	mux.HandleFunc("/", helloHandler)

	// Wrap the ServeMux with a custom handler to catch 404 errors
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the requested path matches any of the defined routes
		_, pattern := mux.Handler(r)
		if pattern == "" {
			// If no route matches, call the custom 404 handler
			notFoundHandler(w, r)
			return
		}
		// Otherwise, serve the request using the default ServeMux
		mux.ServeHTTP(w, r)
	})

	fmt.Printf("Starting server at port %s\n", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		fmt.Println(err)
	}
}
