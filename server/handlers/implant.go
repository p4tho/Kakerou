package handlers

import (
	"fmt"
	"net/http"

	// "server/database"
)

// Retrieve commands from database
func Beacon(w http.ResponseWriter, r *http.Request) {
    fmt.Println("beaconed")
}

// Ping server to check connection
func Ping(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("[*] Pinged from %s\n", r.RemoteAddr)
}
