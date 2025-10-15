package utilsDB

import "myproject/utils"

func villageID(jobID int) int {
	sql := "select villageid from jobs.dbo.jobs where id = "
	sql = sql + utils.Int2string(jobID)
	//GetDBInt(sql) can't import main, so I'll need to dump all of this into
	// its own package.
	// and that's a job for tomorrow!
	// now run this beauty in the DB
	return -1 // for now
}
