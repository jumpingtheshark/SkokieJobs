package main

import (
	"fmt"
	"myproject/Config"
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
		htmlRow = strings.ReplaceAll(htmlRow, "$companyid",
			strconv.Itoa(job.CompanyID))
		htmlRow = strings.ReplaceAll(htmlRow, "$jobtitle", job.JobTitle)
		htmlRow = strings.ReplaceAll(htmlRow, "$jobdescription", job.JobDescription)
		bigInner += htmlRow
	}

	return bigInner
}

func getAllJobs() []entities.Job {
	query := "select * from jobs.dbo.jobs"
	jobs := []entities.Job{}
	rows := utilsDB.GetRows(query)
	defer rows.Close()
	for rows.Next() {
		job := entities.Job{}
		rows.Scan(
			&job.ID,
			&job.CompanyID,
			&job.JobTitle,
			&job.JobDescription,
			&job.Email,
			&job.VillageID)
		jobs = append(jobs, job)
	}
	return jobs
}
