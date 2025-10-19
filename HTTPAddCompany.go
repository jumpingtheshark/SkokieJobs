package main

import (
	"fmt"
	"myproject/utilsDB"
	"net/http"
	"os"
	"strconv"
)

func AddCompany(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		addCompanyGet(w, r)
		return
	} else if r.Method == http.MethodPost {
		addCompanyPost(w, r)
		return
	}

}

func addCompanyPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	companyName := r.FormValue("companyName")
	id := strconv.Itoa(CompanyID())
	insert := fmt.Sprintf(
		"INSERT INTO Companies (CompanyID, CompanyName, villageID) VALUES (%s, '%s', 1)",
		id, companyName,
	)

	utilsDB.InsertUpdate(insert, "", "", "")
	fmt.Fprint(w, companyName+" added")

}

func addCompanyGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	cwd, _ := os.Getwd()
	data, _ := os.ReadFile(cwd + "/UI/AddCompany/AddCompany.html")
	payload := string(data)
	fmt.Fprint(w, payload)

}
