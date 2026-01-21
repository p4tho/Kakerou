package database

import (
	"log"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB *sql.DB
)

/* Initialize database connection */
func DBInit() {
	var err error

	DB, err = sql.Open("sqlite3", "./c2.db")

	// Make sure no error while opening/creating
	if err != nil {
		log.Fatalf("Unable to open SQLite database: %s", err)
	}
	
	// Check connection
	if err = DB.Ping(); err != nil {
		log.Fatalf("Unable to open SQLite database: %s", err)
	}

	// Create tables
	err = execSQLFileSplit("database/sql/tables.sql")
	if err != nil {
		log.Fatalf("Failed to create tables: %s", err)
	}
}