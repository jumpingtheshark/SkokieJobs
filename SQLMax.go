package main

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
