package entities

import (
	"myproject/Config"
	"testing"
)

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

func TestJob_InsertMe(t *testing.T) {

	Config.RunConfig()
	j := Job{}
	j.ID = 1
	j.CompanyID = 999
	j.JobTitle = "test title"
	j.JobDescription = "test description"
	j.Email = "email@test.com"
	j.VillageID = 1
	j.PostingURL = "http://test.com"
	j.PostedBy = "admin"
	j.InsertMe()
}
