package main

import (
	"database/sql"
	"log"
)

func DBPing() {
	dsn := "server=misha-box;database=testdb;trusted_connection=yes;encrypt=true;TrustServerCertificate=true"

	db, err := sql.Open("sqlserver", dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err == nil {
		print("we are connected")
	}
	defer db.Close() // Ensure the connection is closed when the function exits

	rows, err := db.Query("SELECT name FROM students")
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

}
