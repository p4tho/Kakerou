package handlers

import (
	"fmt"
	"net/http"

	"server/database"
)

// Add ping command to database
func PingC2(w http.ResponseWriter, r *http.Request) {
	var cmd database.Command = database.Command {
		Command_id: 0,
		Command: "ping",
		Status: 1,
	}

	database.InsertCommand(&cmd)
}

func Beacon(w http.ResponseWriter, r *http.Request) {
    fmt.Println("beaconed")
}
