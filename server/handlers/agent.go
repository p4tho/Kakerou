package handlers

import (
	"log"
	"net/http"
	"encoding/json"

	"server/database"
)

// Retrieve commands from database (POST)
func Beacon(w http.ResponseWriter, r *http.Request) {
	// Only accpet post requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	// Parse agent's request body
    var req struct {
        Name string `json:"name"`
        UID  int 	`json:"uid"`
    }
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

    log.Printf("%s (UID=%d) beaconed\n",
		req.Name,
		req.UID,
	)
}

// Ping server to check connection (GET)
func Ping(w http.ResponseWriter, r *http.Request) {
    log.Printf("Pinged from %s\n", r.RemoteAddr)
}

// Register agent to database (POST)
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
