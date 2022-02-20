package main

import (
	"testing"
)

var keyWorks = []string{
	"w1",
	"w2",
	"w3",
	"w1",
	"w2",
	"w3",
	"w1",
	"w4",
}

func TestSearchAndCCountResults(t *testing.T) {

	var exp Result
	exp.TotalRecords = 3

	result := Search("w1", keyWorks)

	if result.TotalRecords != exp.TotalRecords {
		t.Errorf("got %d, wanted %d", result.TotalRecords, exp.TotalRecords)
	}
}

func TestSearchAndTotalResults(t *testing.T) {

	var exp Result
	exp.Data = []string{"w1", "w1", "w1"}

	result := Search("w1", keyWorks)

	for i, v := range exp.Data {
		if v != result.Data[i] {
			t.Errorf("got %s, wanted %s", result.Data[i], v)
		}
	}
}

func TestSearchAndTimeResults(t *testing.T) {
	result := Search("w1", keyWorks)

	if result.TimeSearch == "" {
		t.Errorf("got %s, wanted != \"\"", result.TimeSearch)
	}
}

func TestSearchAndErrorResults(t *testing.T) {
	result := Search("w5", keyWorks)

	if result.Message == "" {
		t.Errorf("got %s, wanted != \"\"", result.Message)
	}
}

func TestSearchAndToolsResults(t *testing.T) {
	result := Search("w1", keyWorks)

	if result.Tools == "" {
		t.Errorf("got %s, wanted != \"\"", result.Message)
	}

	result = Search("w5", keyWorks)

	if result.Tools == "" {
		t.Errorf("got %s, wanted != \"\"", result.Message)
	}
}
