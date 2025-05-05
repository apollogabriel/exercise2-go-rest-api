package handlers

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Root Route"))
	fmt.Println("Hello Root Route")

	if r.Method == http.MethodOptions {
		handlePreflight(w, r)
		return
	}

	// Set CORS headers for the main request
	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-Type", "application/json")

	// Example login response
	fmt.Fprintf(w, `{"message":"Login successful"}`)
}

func handlePreflight(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.WriteHeader(http.StatusOK)
}
