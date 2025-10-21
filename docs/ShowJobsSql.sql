select id,
       j.companyID,
       c.CompanyName,
       jobTitle,
       jobDescription,
       email,
       c.villageID,
       v.VillageName
from
    jobs.dbo.jobs j
        inner join jobs.dbo.Companies c on j.companyID=c.CompanyID
        inner join jobs.dbo.Villages v on c.villageID=v.VillageID

where j.id =2


