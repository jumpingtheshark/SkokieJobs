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

func getRows(query string) (rows *sql.Rows, rowCount int) {
	db, _ := sql.Open("sqlserver", Config.dsn)
	defer db.Close()
	rows, _ = db.Query(query)
	i := 0
	for rows.Next() {
		i++
	}
	rows, _ = db.Query(query)
	return rows, i
}

func getDBInt(query string) (i int) {
	db, _ := sql.Open("sqlserver", Config.dsn)
	defer db.Close()
	row := db.QueryRow(query)
	row.Scan(&i)
	return i
}

func getARow(query string) (row *sql.Row) {
	db, _ := sql.Open("sqlserver", Config.dsn)
	fmt.Printf("%T\n", db)
	defer db.Close()
	row = db.QueryRow(query)
	return row
}

func insertUpdate(query string, id string, title string, desc string) {
	db, err := sql.Open("sqlserver", Config.dsn)
	if err != nil {
		log.Fatal(err)
		println("error in db statement")
	}
	defer db.Close()
	exec, err := db.Exec(query)
	if err != nil {
		return
	}
	rows, _ := exec.RowsAffected()
	println(rows)
}
