package main

import (
	"log"
	"net/http"
)

func main() {
	// Load environment variables
	err := loadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize Supabase client
	initSupabase()

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	log.Println("Listening on :8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func loadEnv() error {
	// Implementation for loading environment variables from .env file
	// This function is just a placeholder and should be implemented as needed
	return nil
}

func initSupabase() {
	// Initialize Supabase client with credentials from environment variables
	// This function is just a placeholder and should be implemented as needed
}
