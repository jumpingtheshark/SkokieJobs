package utilsDB

import (
	"fmt"
	"myproject/Config"
	"testing"
)

func TestGetRows(t *testing.T) {

	Config.RunConfig()
	rows := GetRows("select id from jobs.dbo.jobs")

	var s int
	for rows.Next() {
		rows.Scan(&s)
		fmt.Println(s)
	}

}
