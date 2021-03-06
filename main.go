package main

import (
	"encoding/json"
	"log"
	"time"
)

type Result struct {
	Data         []string `json:"data"`
	TotalRecords int      `json:"total_records"`
	Tools        string   `json:"tool"`
	Message      string   `json:"message"`
	TimeSearch   string   `json:"time"`
}

func main() {
	result := Search("w1", []string{"w1", "w2", "w1"})
	data, _ := json.Marshal(result)
	log.Println(string(data))
}

func Search(query string, keyWorks []string) Result {
	var r Result
	var data []string

	start := time.Now()
	for _, v := range keyWorks {
		if query == v {
			data = append(data, query)
		}
	}
	elapsed := time.Since(start)

	r.Tools = "/tools"
	r.TotalRecords = len(data)
	r.TimeSearch = elapsed.String()
	if len(data) > 0 {
		r.Data = data
	} else {
		r.Message = "Error: empty results"
	}

	return r
}
