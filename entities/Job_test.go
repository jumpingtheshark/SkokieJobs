package entities

import "testing"

func TestJob_InsertString(t *testing.T) {
	j := Job{}
	j.ID = 1
	j.CompanyID = 1
	j.JobTitle = "test"
	j.JobDescription = "test"
	j.Email = "test"
	j.VillageID = 1

	s := j.InsertString()
	println(s)
	t.Log(s)
}
