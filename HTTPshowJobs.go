package main

import (
	"fmt"
	"myproject/entities"
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
	//fmt.Fprintf(w, id
	query := "select * from jobs.dbo.jobs where id =" + id

	row := getARow(query)
	data, _ := os.ReadFile("showJob.txt")
	jobTemplate := string(data)

	job := entities.Job{}
	row.Scan(
		&job.ID,
		&job.CompanyID,
		&job.JobTitle,
		&job.JobDescription,
		&job.Email)

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
	fmt.Fprint(w, jobTemplate)

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
