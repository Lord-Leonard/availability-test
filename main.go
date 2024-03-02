package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type data struct {
	TotalRequestCount  int
	FailedRequestCount int
	FailedRequests     []failedRequest
	Errors             []error
}

type failedRequest struct {
	Time         string
	ResponseCode int
}

func main() {
	data := data{
		TotalRequestCount:  0,
		FailedRequestCount: 0,
		FailedRequests:     []failedRequest{},
		Errors:             []error{},
	}

	for {
		res, err := http.Get("https://youtube.com")
		if err != nil {
			data.Errors = append(data.Errors, err)
		}

		data.TotalRequestCount = data.TotalRequestCount + 1

		if res.StatusCode != 200 {
			data.FailedRequestCount = data.FailedRequestCount + 1
			failedRequest := failedRequest{
				time.Now().Format(time.Kitchen),
				res.StatusCode,
			}

			data.FailedRequests = append(data.FailedRequests, failedRequest)
		}

		prettyData, _ := json.MarshalIndent(data, "", "  ")
		fmt.Println(string(prettyData))

		time.Sleep(10 * time.Second)
	}
}
