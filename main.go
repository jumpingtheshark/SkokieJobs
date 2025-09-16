package main

import (
	"database/sql"
	"fmt"
	_ "github.com/microsoft/go-mssqldb" // Blank import for the SQL Server driver
	"log"
)

func main() {
	print("hi")
	//dsn := "sqlserver://username:password@localhost:1433?database=YourDatabaseName"
	//dsn := "server=localhost;database=testdb;trusted_connection=yes"
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

	// Iterate through results
	var shem string
	for rows.Next() {

		rows.Scan(&shem)
		fmt.Printf("Student: %s\n", shem)
	}

}
