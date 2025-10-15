package main

import (
	"fmt"
	"myproject/entities"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func AllJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	query := "select * from jobs.dbo.jobs"
	rows, _ := getRows(query)
	defer rows.Close()
	wd, _ := os.Getwd()
	data, _ := os.ReadFile(wd + "/UI/ShowJobs/inner.html")
	htmlTemplate := string(data)
	bigInner := ""
	var job entities.Job
	for rows.Next() {
		htmlRow := htmlTemplate
		job = entities.Job{}
		rows.Scan(
			&job.ID,
			&job.CompanyID,
			&job.JobTitle,
			&job.JobDescription,
			&job.Email)

		htmlRow = strings.ReplaceAll(htmlRow, "$jobid", strconv.Itoa(job.ID))
		htmlRow = strings.ReplaceAll(htmlRow, "$companyid",
			strconv.Itoa(job.CompanyID))
		htmlRow = strings.ReplaceAll(htmlRow, "$jobtitle", job.JobTitle)
		htmlRow = strings.ReplaceAll(htmlRow, "$jobdescription", job.JobDescription)
		bigInner += htmlRow

	} //for rows.
	data, _ = os.ReadFile(wd + "/UI/ShowJobs/outer.html")
	outerHtmlTemplate := string(data)
	outerHtmlTemplate = strings.ReplaceAll(outerHtmlTemplate, "$inner", bigInner)
	fmt.Fprint(w, outerHtmlTemplate)

}
