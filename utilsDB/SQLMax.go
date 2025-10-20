package utilsDB

import (
	"database/sql"
	"myproject/Config"
)

func JobsID() int {
	sqlstring := "select max(ID) from Jobs"
	i := GetDBInt(sqlstring)
	i++
	return i
}

func CompanyID() int {
	sqlstring := "select max(companyid) from Companies"
	i := GetDBInt(sqlstring)
	i++
	return i
}

func GetMaxJobID() int {
	sqlstring := "select max(ID) from jobs.dbo.jobs"
	db, _ := sql.Open("sqlserver", Config.Config.Dsn)
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
