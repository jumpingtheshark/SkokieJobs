package main

import (
	_ "github.com/microsoft/go-mssqldb" // Blank import for the SQL Server driver
	"myproject/Config"
	"myproject/utilsDB"
)

func main() {

	/*
		update
		10-15-2025
		5 pm

		inserting works.
		next step, make  a function that looks up village id based on job id

	*/

	Config.RunConfig()
	utilsDB.DBPing()
	RunServer()
}
