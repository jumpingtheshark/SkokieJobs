package utilsDB

import (
	"database/sql"
	"fmt"
	"log"
	"myproject/Config"
)

//var dsn = "server=misha-box;database=testdb;trusted_connection=yes;encrypt=true;TrustServerCertificate=true"

func DBPing() {

	db, err := sql.Open("sqlserver", Config.Config.Dsn)

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

func GetRows(query string) (rows *sql.Rows, rowCount int) {
	db, _ := sql.Open("sqlserver", Config.Config.Dsn)
	defer db.Close()

	//fmt.Printf("%T\n", db)
	rows, _ = db.Query(query)
	i := 0
	for rows.Next() {
		i++
	}
	rows, _ = db.Query(query)
	return rows, i
}

func GetDBInt(query string) (i int) {
	db, _ := sql.Open("sqlserver", Config.Config.Dsn)
	defer db.Close()
	row := db.QueryRow(query)
	row.Scan(&i)
	return i
}

func InsertUpdate2(sqlCommand string) int64 {
	db, err := sql.Open("sqlserver", Config.Config.Dsn)
	if err != nil {
		log.Fatal(err)
		println("db object failed")
	}
	defer db.Close()
	exec, err := db.Exec(sqlCommand)
	if err != nil {
		return -1
	}
	rows, _ := exec.RowsAffected()
	println(rows)
	return rows
}
func InsertUpdate(query string, id string, title string, desc string) {
	db, err := sql.Open("sqlserver", Config.Config.Dsn)
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
func GetARow(query string) (row *sql.Row) {
	db, _ := sql.Open("sqlserver", Config.Config.Dsn)
	fmt.Printf("%T\n", db)
	defer db.Close()
	row = db.QueryRow(query)
	return row
}
