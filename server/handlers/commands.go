package handlers

import (
	"net/http"

	"server/database"
)

// Add ping command to queue
func PingC2(w http.ResponseWriter, r *http.Request) {
	var cmd database.Command = database.Command {
		Command_id: 0,
		Command: "ping",
		Status: 1,
	}

	database.InsertCommand(&cmd)
}