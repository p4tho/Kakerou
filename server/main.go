package main

import (
	"server/c2"
	"server/database"
)

func main() {
	// Use config file to determine settings
	var port string = "8080"

	// Create database connection
	database.DBInit()
	defer database.DB.Close()

	// Run C2 server (run as goroutine?)
	c2.HttpInit(port)
}