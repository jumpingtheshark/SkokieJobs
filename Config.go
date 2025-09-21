package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var CFG map[string]string

type Configuration struct {
	dsn    string
	curdir string
}

var Config Configuration

func RunConfig() {
	// Read file
	data, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	// Map of strings
	if err := json.Unmarshal(data, &CFG); err != nil {
		panic(err)
	}

	CFG["curdir"], _ = os.Getwd()

	Config = Configuration{
		dsn:    CFG["dsn"],
		curdir: CFG["curdir"],
	}
	// Access values directly
	fmt.Println("DSN:", CFG["dsn"])

}
