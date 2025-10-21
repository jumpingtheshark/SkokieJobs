package main

import (
	"fmt"
	"testing"
)

func TestFillData(t *testing.T) {

	job := FillData("2")
	fmt.Println(job.JobTitle)

}
