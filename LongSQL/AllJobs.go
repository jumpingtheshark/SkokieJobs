package LongSQL

func AllJobs() string {
	s := `select id,
		c.CompanyName,
		jobTitle,
		jobDescription,
		v.VillageName as villageName,
		datePosted
		from
		jobs.dbo.jobs j
left join jobs.dbo.Villages v on j.villageID=v.VillageID
left join jobs.dbo.Companies c on j.companyID=c.CompanyID
order by j.id desc
`
	return s
}
