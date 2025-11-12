package entities

import (
	"database/sql"
	"fmt"
	"myproject/utils"
	"myproject/utilsDB"
	"strconv"
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

func (j *Job) LoadMe() {
	query := j.SelectString()
	dbRow := utilsDB.GetRows(query)
	j.SetValues(dbRow)

}

func (j *Job) SetValues(dbRow *sql.Rows) {
	for dbRow.Next() {
		dbRow.Scan(
			&j.CompanyID,
			&j.CompanyName,
			&j.JobTitle,
			&j.JobDescription,
			&j.Email,
			&j.VillageID,
			&j.VillageName,
			&j.DatePosted,
			&j.PostingURL,
		)

	}

	j.DatePostedString = utils.Date2string(j.DatePosted)

}
func (j *Job) SelectString() string {
	query := `select
       j.companyID,
       c.CompanyName,
       jobTitle,
       jobDescription,
       email,
       j.villageID,
       v.VillageName,
      j.datePosted,
      j.postingURL
       from  
           jobs.dbo.jobs j 
           inner join 
           jobs.dbo.Companies c 
           on j.companyID=c.CompanyID
           inner join jobs.dbo.Villages v
           on c.VillageID= v.VillageID
           where j.id =`
	query = query + strconv.Itoa(j.ID)

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
                      PostingURL,
                      isActive)
VALUES (%d, 
        %d, 
        '%s', 
        '%s', 
        '%s',
        %d,
        '%s',
        '%s',
        '%s',
        %d);`

	insert = fmt.Sprintf(insert,
		j.ID,
		j.CompanyID,
		j.JobTitle,
		j.JobDescription,
		j.Email,
		j.VillageID,
		j.DatePostedString,
		j.PostedBy,
		j.PostingURL,
		1,
	)

	return insert
}
