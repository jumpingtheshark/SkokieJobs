package main

import (
	"database/sql"
	"fmt"
	"log"
)

//var dsn = "server=misha-box;database=testdb;trusted_connection=yes;encrypt=true;TrustServerCertificate=true"

func DBPing() {

	db, err := sql.Open("sqlserver", Config.dsn)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err == nil {
		print("we are connected ")
	}
	defer db.Close() // Ensure the connection is closed when the function exits

	rows, err := db.Query("SELECT top 1 *  FROM jobs")
	if err != nil {
		log.Fatal("Query failed:", err)
	}
	defer rows.Close()

}

func getRows(query string) (rows *sql.Rows) {
	db, _ := sql.Open("sqlserver", Config.dsn)
	defer db.Close()
	rows, _ = db.Query(query)
	return rows
}

func getARow(query string) (row *sql.Row) {
	db, _ := sql.Open("sqlserver", Config.dsn)
	fmt.Printf("%T\n", db)
	defer db.Close()
	row = db.QueryRow(query)
	return row
}
