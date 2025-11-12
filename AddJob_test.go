package main

import (
	"fmt"
	"myproject/Config"
	"testing"
)

func TestTemplate(t *testing.T) {

	s := Template()
	fmt.Println(s)

}

func TestTryDate(t *testing.T) {
	Config.RunConfig()
	TryDate()
}
