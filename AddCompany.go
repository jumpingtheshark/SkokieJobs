package main

import (
	"fmt"
	"myproject/UI/UIComponents"
	"myproject/entities"
	"myproject/utilsDB"
	"net/http"
	"os"
	"strconv"
	"strings"
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
	c := entities.Company{}

	r.ParseForm()
	c.CompanyName = r.FormValue("companyName")
	c.Id = strconv.Itoa(utilsDB.CompanyID())
	c.VillageID = r.FormValue("villages")
	c.AddressLine1 = r.FormValue("addressLine1")
	c.AddressLine2 = r.FormValue("addressLine2")
	c.Zip = r.FormValue("zip")
	insert := fmt.Sprintf(
		`INSERT INTO Companies 
    		(CompanyID, 
    		 CompanyName, 
    		 villageID,
    		 AddressLine1,
    		 AddressLine2,
    		 zip) 
			VALUES 
			(%s, '%s', %s, '%s', '%s', '%s')`,
		c.Id,
		c.CompanyName,
		c.VillageID,
		c.AddressLine1,
		c.AddressLine2,
		c.Zip)

	utilsDB.InsertUpdate(insert, "", "", "")
	fmt.Fprint(w, c.CompanyName+" added")

}

func addCompanyGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	cwd, _ := os.Getwd()
	data, _ := os.ReadFile(cwd + "/UI/AddCompany/AddCompany.html")
	payload := string(data)
	villagesDDL := UIComponents.VillagesDDL()
	payload = strings.ReplaceAll(payload, "$villagesDDL", villagesDDL)
	fmt.Fprint(w, payload)
}
