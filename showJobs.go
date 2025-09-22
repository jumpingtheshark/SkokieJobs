package main

import (
	"fmt"
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

	//fmt.Fprintf(w, "<html><body>")

	var dbid int
	var companyID int
	var jobTitle string
	var jobDescription string
	var email string
	row.Scan(&dbid, &companyID, &jobTitle, &jobDescription, &email)
	jobTemplate = strings.ReplaceAll(jobTemplate, "$title", jobTitle)
	jobTemplate = strings.ReplaceAll(jobTemplate, "$description", jobDescription)
	jobTemplate = strings.ReplaceAll(jobTemplate, "$jobID", strconv.Itoa(dbid))
	fmt.Fprintf(w, jobTemplate)

	//fmt.Fprintf(w, strconv.Itoa(dbid)+"<br>")
	//fmt.Fprintf(w, jobTitle+"<br>")
	//fmt.Fprintf(w, jobDescription+"<br>")
	//fmt.Fprintf(w, email)

	//fmt.Fprintf(w, "</body></html>")

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
