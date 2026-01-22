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
	var err error

	err = json.NewDecoder(r.Body).Decode(&agent)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	uid, err := database.InsertAgent(&agent)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Return uid in response
	response := struct {
		UID  int64  `json:"uid"`
	}{
		UID:  uid,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
