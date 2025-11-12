package entities

import (
	"fmt"
	"myproject/utils"
	"myproject/utilsDB"
	"time"
)

type Job struct {
	ID               int
	CompanyID        int
	CompanyName      string
	JobTitle         string
	JobDescription   string
	Email            string
	VillageID        int
	VillageName      string
	DatePosted       time.Time
	DateLastEdited   time.Time
	DatePostedString string
	PostedBy         string
	LastEditedBy     string
	PostingURL       string
	IsActive         bool
}

func (j *Job) InsertMe() {
	insertString := j.InsertString()
	utilsDB.InsertUpdate2(insertString)

}

func (j *Job) LoadMe() {}

func (j *Job) SelectString() string {
	query := `select id,
       companyID,
       jobTitle,
       jobDescription,
       email,
       villageID
       from 
           jobs.dbo.jobs where id =`
	return query
}

func (j *Job) InsertString() string {

	j.DatePosted = utils.TodayDate()
	j.DatePostedString = utils.Date2string(j.DatePosted)

	insert := `
INSERT INTO dbo.Jobs (
                      ID, 
                      CompanyID, 
                      JobTitle, 
                      JobDescription, 
                      Email,
                      villageID,
                      DatePosted,
                      postedBy,
                      PostingURL)
VALUES (%d, 
        %d, 
        '%s', 
        '%s', 
        '%s',
        %d,
        '%s',
        '%s',
        '%s');`

	insert = fmt.Sprintf(insert,
		j.ID,
		j.CompanyID,
		j.JobTitle,
		j.JobDescription,
		j.Email,
		j.VillageID,
		j.DatePostedString,
		j.PostedBy,
		j.PostingURL)

	return insert
}
