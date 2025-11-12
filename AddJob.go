package main

import (
	"fmt"
	"html"
	"myproject/Config"
	"myproject/entities"
	_ "myproject/entities"
	"myproject/utils"
	_ "myproject/utils"
	"myproject/utilsDB"
	"net/http"
	"strings"
	"time"
)

// Route handler: dispatch GET vs POST
func AddJob(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		AddJobGet(w, r)
	case http.MethodPost:
		AddJobPost(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
func AddJobGet(w http.ResponseWriter, r *http.Request) {
	rows := utilsDB.GetRows(`
		SELECT CompanyID, 
		       CompanyName
		FROM dbo.Companies
		ORDER BY CompanyName`)

	var options strings.Builder

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			http.Error(w,
				"scan error: "+err.Error(),
				http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(&options, `<option value="%d">%s</option>`,
			id, html.EscapeString(name))
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "rows error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	page := Template()
	page = strings.ReplaceAll(page, "$options", options.String())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(page))
}

func Template() string {
	bf := Config.BaseDirPath
	path := bf + "ui\\addJob\\addJob.html"
	s := utils.LoadFile(path)
	return s
}

// POST!!!!
// POST: process the form and insert a job
func AddJobPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form",
			http.StatusBadRequest)
		return
	}
	newJob := entities.Job{}

	newJob.CompanyID = utils.String2int(r.FormValue("companyID"))
	newJob.VillageID = utilsDB.VillageID(newJob.CompanyID) //we'll just keep in Skokie for now I guess

	newJob.JobTitle = r.FormValue("jobTitle")

	newJob.JobDescription = r.FormValue("jobDescription")
	newJob.ID = utilsDB.GetMaxJobID() + 1
	newJob.Email = r.FormValue("email")
	newJob.PostingURL = r.FormValue("postingURL")
	newJob.PostedBy = "admin"
	newJob.InsertMe()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(
		utils.Int2string(newJob.ID) + "Job added"))

	//	insertUpdate(insert, jobID, companyID, title, desc)

}

func TryDate() {
	j := entities.Job{}
	today := time.Date(
		time.Now().Year(),
		time.Now().Month(),
		time.Now().Day(),
		0,
		0,
		0,
		0,
		time.UTC)

	j.DatePosted = today
	postedString := j.DatePosted.Format("2006-01-02")
	fmt.Println(j.DatePosted)
	sql := fmt.Sprintf(`
		INSERT INTO dbo.jobs (
		                  id, 
		                  datePosted
		                  )
		VALUES (%s, '%s')
	`,
		"999",
		postedString)

	fmt.Println(sql)
	utilsDB.InsertUpdate2(sql)

}
