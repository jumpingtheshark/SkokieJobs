package main

import (
	"myproject/Config"
	"myproject/entities"
	"myproject/utils"
	"net/http"
	"strings"
)

func NewJob(
	w http.ResponseWriter,
	r *http.Request) {
	switch r.Method {
	case
		http.MethodGet:
		NewJobGet(w, r)
	case
		http.MethodPost:
		NewJobPost(w, r)
	default:
		http.Error(w, "method not allowed",
			http.StatusMethodNotAllowed)
	}
}

func NewJobGet(
	w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	parts := strings.Split(r.URL.Path, "/")
	length := len(parts)
	id := parts[length-1]
	company := entities.Company{}
	company.Id = id
	company.LoadMe()
	template := GetJobTemplate()
	template = strings.ReplaceAll(template, "$companyid", company.Id)
	template = strings.ReplaceAll(template, "$companyname", company.CompanyName)
	w.Write([]byte(template))

}

func GetJobTemplate() string {
	htmlPath := Config.UIPaths.NewJob
	template := utils.LoadFile(htmlPath)
	return template

}
