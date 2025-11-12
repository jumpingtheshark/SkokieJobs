package main

import (
	"fmt"
	"myproject/Config"
	"myproject/entities"
	"myproject/utils"
	"myproject/utilsDB"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func ShowJob(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	parts := strings.Split(r.URL.Path, "/")
	length := len(parts)
	id := parts[length-1]

	job := FillData(id)
	output := BuildOutputString2(job)
	fmt.Fprint(w, output)

}
func BuildOutputString2(job entities.Job) string {
	jobTemplate := GetTemplate()
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$title", job.JobTitle)
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$description", job.JobDescription)
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$jobID", strconv.Itoa(job.ID))
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$email", job.Email)
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$companyID", strconv.Itoa(job.CompanyID))
	return jobTemplate
}
func GetTemplate() string {
	htmlPath := Config.UIPaths.JobDetails
	template := utils.LoadFile(htmlPath)
	return template

}

func BuildOutputString(job entities.Job) string {

	data, _ := os.ReadFile("showJob.txt")
	jobTemplate := string(data)

	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$title", job.JobTitle)
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$description", job.JobDescription)
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$jobID", strconv.Itoa(job.ID))
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$email", job.Email)
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$companyID", strconv.Itoa(job.CompanyID))
	return jobTemplate
}
func FillData(id string) entities.Job {

	job := entities.Job{}
	query := `select id,
       companyID,
       jobTitle,
       jobDescription,
       email,
       villageID
       from 
           jobs.dbo.jobs where id =`
	query = query + id
	row := utilsDB.GetRows(query)
	for row.Next() {
		row.Scan(
			&job.ID,
			&job.CompanyID,
			&job.JobTitle,
			&job.JobDescription,
			&job.Email,
			&job.VillageID)

	}
	return job
}

/*
	fmt.Fprintf(w, id)
	fmt.Println(length)
	fmt.Println(parts)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.URL.Host)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RequestURI())
*/
