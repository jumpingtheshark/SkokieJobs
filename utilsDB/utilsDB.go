package utilsDB

import (
	"myproject/utils"
)

func VillageID(companyID int) int {
	sql := "select villageid from jobs.dbo.Companies where CompanyID= " + utils.Int2string(companyID)
	retVal := -1
	retVal = GetDBInt(sql) //can't import main, so I'll need to dump all of this into
	// its own package.
	// and that's a job for tomorrow!
	// now run this beauty in the DB
	return retVal // for now
}
