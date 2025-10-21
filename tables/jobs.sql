CREATE TABLE [dbo].[jobs](
    [id] [int] NULL,
    [companyID] [int] NULL,
    [jobTitle] [varchar](256) NULL,
    [jobDescription] [text] NULL,
    [email] [varchar](256) NULL,
    [villageID] [int] NULL,
    [datePosted] [date] NULL,
    [dateLastEdited] [date] NULL,
    [postedBy] [varchar](max) NULL,
    [lastEditedBy] [varchar](max) NULL,
    [postingURL] [varchar](max) NULL
    )