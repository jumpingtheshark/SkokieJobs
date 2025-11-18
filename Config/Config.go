package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

var UIPaths = struct {
	JobDetails     string
	CompaniesInner string
	CompaniesOuter string
	NewJob         string
}{
	JobDetails:     BaseDirPath + "UI\\JobDetails\\JobDetails.html",
	CompaniesInner: BaseDirPath + "UI\\Companies\\inner.html",
	CompaniesOuter: BaseDirPath + "UI\\Companies\\outer.html",
	NewJob:         BaseDirPath + "UI\\NewJob\\newJob.html",
}

var HTTPPaths = struct {
	JobDetails   string
	AddCompany   string
	AddJob       string
	AllJobs      string
	AllCompanies string
	NewJob       string
	NewJobPost   string
}{
	AllJobs:      Config.Server + "/alljobs",
	AddCompany:   Config.Server + "/addCompany",
	AddJob:       Config.Server + "/addJob",
	JobDetails:   Config.Server + "/job/",
	AllCompanies: Config.Server + "/AllCompanies",
	NewJob:       Config.Server + "/newJob/",
	NewJobPost:   Config.Server + "/newJobPost",
}
var CFG map[string]string

var ConfigPath = BaseDirPath + "\\config.json"
var BaseDirPath = "c:\\Users\\mrubi\\GolandProjects\\SkokieJobs\\"
var CurDir, _ = os.Getwd()

type Configuration struct {
	Dsn    string
	Curdir string
	Server string
}

var Config Configuration

func RunConfig() {
	// Read file

	data, err := os.ReadFile(ConfigPath)
	if err != nil {
		panic(err)
	}

	// Map of strings
	if err := json.Unmarshal(data, &CFG); err != nil {
		panic(err)
	}

	CFG["curdir"], _ = os.Getwd()

	Config = Configuration{
		Dsn:    CFG["dsn"],
		Curdir: CFG["curdir"],
		Server: CFG["server"],
	}
	// Access values directly
	fmt.Println("DSN:", CFG["dsn"])

}
