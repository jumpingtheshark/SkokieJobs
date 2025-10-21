package Config

import (
	"encoding/json"
	"fmt"
	"os"
)

var CFG map[string]string

var ConfigPath = BaseDirPath + "\\config.json"
var BaseDirPath = "c:\\Users\\mrubi\\GolandProjects\\SkokieJobs\\"
var CurDir, _ = os.Getwd()

type Configuration struct {
	Dsn    string
	Curdir string
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
	}
	// Access values directly
	fmt.Println("DSN:", CFG["dsn"])

}
