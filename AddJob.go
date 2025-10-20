package main

import (
	"fmt"
	"html"
	"myproject/entities"
	_ "myproject/entities"
	"myproject/utils"
	_ "myproject/utils"
	"myproject/utilsDB"
	"net/http"
	"strings"
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
		SELECT CompanyID, CompanyName
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
		fmt.Fprintf(&options, `<option value="%d">%s</option>`, id, html.EscapeString(name))
	}
	if err := rows.Err(); err != nil {
		http.Error(w, "rows error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	page := `<!doctype html>
<html>
  <body>
    <h2>Add Job</h2>
    <form method="post" action="/addJob">
      <label>Company:
        <select name="companyID" required>
          ` + options.String() + `
        </select>
      </label><br/><br/>
      <label>Job Title: <input name="jobTitle" required></label><br/><br/>
      <label>Email: <input name="email" type="email"></label><br/><br/>
      <label>Description:<br/>
        <textarea name="jobDescription" rows="6" cols="60"></textarea>
      </label><br/><br/>
      <button type="submit">Add</button>
    </form>
    <p><a href="/alljobs">Back to all jobs</a></p>
  </body>
</html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(page))
}

// POST: process the form and insert a job
func AddJobPost(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "bad form",
			http.StatusBadRequest)
		return
	}
	newJob := entities.Job{}

	newJob.CompanyID = utils.String2int(r.FormValue("companyID"))

	newJob.JobTitle = r.FormValue("jobTitle")

	newJob.JobDescription = r.FormValue("jobDescription")
	newJob.ID = utilsDB.GetMaxJobID() + 1
	newJob.Email = r.FormValue("email")
	newJob.VillageID = utilsDB.VillageID(newJob.CompanyID) //we'll just keep in Skokie for now I guess
	insertString := newJob.InsertString()
	inserted := utilsDB.InsertUpdate2(insertString)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(
		utils.Bigint2string(inserted) + "Job added"))

	//	insertUpdate(insert, jobID, companyID, title, desc)

}
