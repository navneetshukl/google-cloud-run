package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func GetRouter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Important: We use 0.0.0.0 to ensure the server is accessible 
	// from outside the container.
	address := fmt.Sprintf("0.0.0.0:%s", port)
	
	http.HandleFunc("/", GetRouter)

	log.Printf("Listening on port %s...", port)
	
	// Use Fatal so if the server fails to start, the logs tell us why immediately
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Error starting the server: %v", err)
	}
}