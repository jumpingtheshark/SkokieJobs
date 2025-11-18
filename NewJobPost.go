package main

import (
	"myproject/entities"
	"myproject/utils"
	"myproject/utilsDB"
	"net/http"
)

func NewJobPost(
	w http.ResponseWriter,
	r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form",
			http.StatusBadRequest)
		return
	}

	newJob := entities.Job{}

	newJob.CompanyID = utils.String2int(r.FormValue("companyID"))
	newJob.VillageID = utilsDB.VillageID(newJob.CompanyID)
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

}
