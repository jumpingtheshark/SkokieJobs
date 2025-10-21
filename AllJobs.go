package main

import (
	"fmt"
	"myproject/Config"
	"myproject/LongSQL"
	"myproject/entities"
	"myproject/utilsDB"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func AllJobs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	jobs := getAllJobs()
	bigInner := BuildInnerRows(jobs)
	totalPage := BuildOuterRows(bigInner)
	fmt.Fprint(w, totalPage)

}
func BuildOuterRows(bigInner string) string {
	wd := Config.BaseDirPath
	data, _ := os.ReadFile(wd + "/UI/ShowJobs/outer.html")
	outerHtmlTemplate := string(data)
	outerHtmlTemplate = strings.ReplaceAll(outerHtmlTemplate, "$inner", bigInner)
	return outerHtmlTemplate
}

func BuildInnerRows(jobs []entities.Job) string {
	wd := Config.BaseDirPath
	data, _ := os.ReadFile(wd + "/UI/ShowJobs/inner.html")
	htmlTemplate := string(data)
	bigInner := ""
	for _, job := range jobs {
		htmlRow := htmlTemplate
		htmlRow = strings.ReplaceAll(htmlRow, "$jobid", strconv.Itoa(job.ID))
		htmlRow = strings.ReplaceAll(htmlRow, "$companyName", job.CompanyName)
		htmlRow = strings.ReplaceAll(htmlRow, "$jobTitle", job.JobTitle)
		htmlRow = strings.ReplaceAll(htmlRow, "$jobDescription", job.JobDescription)
		htmlRow = strings.ReplaceAll(htmlRow, "$villageName", job.VillageName)
		htmlRow = strings.ReplaceAll(htmlRow, "$postedDate", job.DatePosted.String())

		bigInner += htmlRow
	}

	return bigInner
}

func getAllJobs() []entities.Job {
	query := LongSQL.AllJobs()
	jobs := []entities.Job{}
	rows := utilsDB.GetRows(query)
	defer rows.Close()
	for rows.Next() {
		job := entities.Job{}
		rows.Scan(
			&job.ID,
			&job.CompanyName,
			&job.JobTitle,
			&job.JobDescription,
			&job.VillageName,
			&job.DatePosted)
		jobs = append(jobs, job)
	}
	return jobs
}
