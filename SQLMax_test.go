package main

import "testing"

func TestJobsID(t *testing.T) {

	i := JobsID()
	println(i)
	t.Log(i)

}

func TestGetMaxJobID(t *testing.T) {
	i := getMaxJobID()
	println(i)
}
