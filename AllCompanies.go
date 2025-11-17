package main

import (
	"fmt"
	"myproject/Config"
	"myproject/LongSQL"
	"myproject/entities"
	"myproject/utilsDB"
	"net/http"
	"os"
	"strings"
)

func AllCompanies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	companies := getAllCompanies()
	bigInner := BuildInnerRowsCompanies(companies)
	totalPage := BuildOuterRowsCompanies(bigInner)
	fmt.Fprint(w, totalPage)

}

func BuildOuterRowsCompanies(bigInner string) string {
	htmlPath := Config.UIPaths.CompaniesOuter
	data, _ := os.ReadFile(htmlPath)
	outerHtmlTemplate := string(data)
	outerHtmlTemplate = strings.ReplaceAll(outerHtmlTemplate, "$inner", bigInner)
	return outerHtmlTemplate

}
func BuildInnerRowsCompanies(companies []entities.Company) string {
	htmlPath := Config.UIPaths.CompaniesInner
	data, _ := os.ReadFile(htmlPath)
	htmlTemplate := string(data)
	bigInner := ""
	njURL1 := Config.HTTPPaths.NewJob
	nj := ""
	for _, company := range companies {
		htmlRow := htmlTemplate
		htmlRow = strings.ReplaceAll(htmlRow, "$companyid", company.Id)
		htmlRow = strings.ReplaceAll(htmlRow, "$companyname", company.CompanyName)
		htmlRow = strings.ReplaceAll(htmlRow, "$village", company.VillageName)
		nj = njURL1 + string(company.Id)
		htmlRow = strings.ReplaceAll(htmlRow, "$newjob", nj)
		bigInner += htmlRow
	}

	return bigInner
}

func getAllCompanies() []entities.Company {
	query := LongSQL.AllCompanies()
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
