package entities

import (
	"myproject/Config"
	"myproject/utilsDB"
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
	j.ID = utilsDB.JobsID()
	j.CompanyID = 3
	j.JobTitle = "test title"
	j.JobDescription = "test description"
	j.Email = "email@test.com"
	j.VillageID = 1
	j.PostingURL = "http://test.com"
	j.PostedBy = "admin"
	j.InsertMe()
}

func TestJob_LoadMe(t *testing.T) {
	j := Job{}
	j.ID = 1000
	j.LoadMe()
	println(j.JobTitle)

}
