package entities

import (
	"fmt"
	"time"
)

type Job struct {
	ID             int
	CompanyID      int
	CompanyName    string
	JobTitle       string
	JobDescription string
	Email          string
	VillageID      int
	VillageName    string
	DatePosted     time.Time
	DateLastEdited time.Time
	PostedBy       string
	LastEditedBy   string
	PostingURL     string
}

// to do - change this into full out insert and then keep updating 3 by 3.
func (j *Job) InsertString() string {

	insert := `
INSERT INTO dbo.Jobs (
                      ID, 
                      CompanyID, 
                      JobTitle, 
                      JobDescription, 
                      Email,
                      villageID)
VALUES (%d, 
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
		j.VillageID)

	return insert
}
