package entities

import "fmt"

type Job struct {
	ID             int
	CompanyID      int
	JobTitle       string
	JobDescription string
	Email          string
	VillageID      int
}

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
