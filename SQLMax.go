package main

import "database/sql"

func JobsID() int {
	sqlstring := "select max(ID) from Jobs"
	i := getDBInt(sqlstring)
	i++
	return i
}

func CompanyID() int {
	sqlstring := "select max(companyid) from Companies"
	i := getDBInt(sqlstring)
	i++
	return i
}

func getMaxJobID() int {
	sqlstring := "select max(ID) from jobs.dbo.jobs"
	db, _ := sql.Open("sqlserver", Config.dsn)
	defer db.Close()
	i := 0
	row, err := db.Query(sqlstring)
	if err != nil {
		panic(err)
	}
	for row.Next() {
		row.Scan(&i)
	}

	return i
}
