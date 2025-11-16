package main

import (
	"fmt"
	"myproject/LongSQL"
	"myproject/entities"
	"myproject/utilsDB"
	"net/http"
)

func AllCompanies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	companies := getAllCompanies()
	bigInner := BuildInnerRowsCompanies(companies)
	totalPage := BuildOuterRowsCompanies(bigInner)
	fmt.Fprint(w, totalPage)

}

func BuildInnerRowsCompanies(companies []entities.Company) string {
	return ""
}

func BuildOuterRowsCompanies(bigInner string) string {
	return ""
}
func getAllCompanies() []entities.Company {
	query := LongSQL.AllJobs()
	companies := []entities.Company{}
	rows := utilsDB.GetRows(query)
	defer rows.Close()
	for rows.Next() {
		company := entities.Company{}
		rows.Scan(
			&company.Id,
			&company.CompanyName,
			&company.VillageName,
		)
		companies = append(companies, company)

	}
	return companies
}
