package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"

	"server/database"
)

// Retrieve commands from database
func Beacon(w http.ResponseWriter, r *http.Request) {
    fmt.Println("beaconed")
}

// Ping server to check connection
func Ping(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("[*] Pinged from %s\n", r.RemoteAddr)
}

// Register agent to database
func Register(w http.ResponseWriter, r *http.Request) {
	// Only accept post requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	// Get agent data (just name) from request
	var agent database.Agent

	err := json.NewDecoder(r.Body).Decode(&agent)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	database.InsertAgent(&agent)
}
