package main

import (
	"fmt"

	"net/http"
)

func RunServer() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/job/", ShowJob)
	http.HandleFunc("/alljobs", AllJobs)
	http.HandleFunc("/", hello)
	http.HandleFunc("/addCompany", AddCompany)
	http.HandleFunc("/addJob", AddJob)

	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "I am online, shalom y'all ")
}
