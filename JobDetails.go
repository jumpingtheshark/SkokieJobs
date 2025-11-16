package main

import (
	"fmt"
	"myproject/Config"
	"myproject/entities"
	"myproject/utils"
	"net/http"
	"strconv"
	"strings"
)

func ShowJob(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	parts := strings.Split(r.URL.Path, "/")
	length := len(parts)
	id := parts[length-1]

	//job := FillData(id)
	job := entities.Job{}
	job.ID = utils.String2int(id)
	job.LoadMe()
	output := BuildOutputString2(job)
	fmt.Fprint(w, output)

}
func BuildOutputString2(job entities.Job) string {
	jobTemplate := GetTemplate()

	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$alljobs", Config.HTTPPaths.AllJobs)
	//title
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$title", job.JobTitle)

	//jobID
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$jobID", strconv.Itoa(job.ID))

	//companyName
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$companyName", job.CompanyName)

	//villageName
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$villageName", job.VillageName)

	//postedDate
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$postedDate", job.DatePostedString)

	//description
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$description", job.JobDescription)

	//email
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$email", job.Email)

	//postingURL
	jobTemplate = strings.ReplaceAll(jobTemplate,
		"$postingURL", job.PostingURL)

	return jobTemplate
}

func GetTemplate() string {
	htmlPath := Config.UIPaths.JobDetails
	template := utils.LoadFile(htmlPath)
	return template

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
