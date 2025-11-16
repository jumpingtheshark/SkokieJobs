package entities

import (
	"database/sql"
	"myproject/utilsDB"
)

type Company struct {
	Id           string
	CompanyName  string
	VillageID    string
	VillageName  string
	AddressLine1 string
	AddressLine2 string
	Zip          string
}

func (c *Company) LooadMe() {
	query := c.SelectString()
	dbRow := utilsDB.GetRows(query)
	c.SetValues(dbRow)

}

func (c *Company) SetValues(dbRow *sql.Rows) {
	for dbRow.Next() {
		dbRow.Scan(
			&c.Id,
			&c.CompanyName,
			&c.VillageName,
		)

	}
}
func (c *Company) SelectString() string {
	query := `
select 
companyID,
CompanyName,
VillageName
from 
companies c
inner join 
villages v on 
c.villageID=v.VillageID
`
	return query
}
