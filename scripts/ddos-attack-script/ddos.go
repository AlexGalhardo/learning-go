package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"
)

var listEndpointsToAttack = []string{
	"https://jsonplaceholder.typicode.com/users/1",
	"https://jsonplaceholder.typicode.com/users/2",
	"https://jsonplaceholder.typicode.com/users/3",
	"https://jsonplaceholder.typicode.com/posts/1",
	"https://jsonplaceholder.typicode.com/posts/2",
	"https://jsonplaceholder.typicode.com/posts/3",
	"https://jsonplaceholder.typicode.com/comments/1",
	"https://jsonplaceholder.typicode.com/comments/2",
	"https://jsonplaceholder.typicode.com/comments/3",
	"https://jsonplaceholder.typicode.com/todos/1",
	"https://jsonplaceholder.typicode.com/todos/2",
	"https://jsonplaceholder.typicode.com/todos/3",
}

type Request struct {
	Endpoint               string      `json:"endpoint"`
	HTTPStatusCodeResponse int         `json:"http_status_code_response"`
	Response               interface{} `json:"response"`
	ResponseSuccess        bool        `json:"response_success"`
}

type WorkerJobStatistics struct {
	WorkerID                     int         `json:"worker_id"`
	WorkerExecutionTime          string      `json:"worker_execution_time"`
	TotalRequestsMade            int         `json:"total_requests_made"`
	TotalRequestsTimeout         int         `json:"total_requests_timeout"`
	TotalRequestsHTTPStatusCode200 int       `json:"total_requests_http_status_code_200"`
	TotalRequestsFail            int         `json:"total_requests_fail"`
	Responses                    []Request   `json:"responses"`
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	start := time.Now()

	fmt.Printf("Worker ID: %d is online\n", id)

	totalRequestsMade := 0

	workerJobStatistics := WorkerJobStatistics{
		WorkerID: id,
	}

	for i := 0; i < 20; i++ {
		for _, endpoint := range listEndpointsToAttack {
			totalRequestsMade++

			request := Request{
				Endpoint: endpoint,
			}

			fmt.Printf("Worker ID: %d => Processing REQUEST GET: %s\n", id, endpoint)

			resp, err := http.Get(endpoint)
			if err != nil {
				workerJobStatistics.TotalRequestsTimeout++
				log.Printf("Error fetching URL %s: %v", endpoint, err)
				continue
			}

			request.HTTPStatusCodeResponse = resp.StatusCode
			if resp.StatusCode == 200 {
				workerJobStatistics.TotalRequestsHTTPStatusCode200++
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				workerJobStatistics.TotalRequestsFail++
				log.Printf("Error reading response body: %v", err)
				continue
			}

			err = json.Unmarshal(body, &request.Response)
			if err != nil {
				workerJobStatistics.TotalRequestsFail++
				log.Printf("Error unmarshalling response: %v", err)
				continue
			}

			request.ResponseSuccess = true
			workerJobStatistics.Responses = append(workerJobStatistics.Responses, request)

			err = writeFile(fmt.Sprintf("./responses/responses-from-worker-id-%d.json", id), workerJobStatistics)
			if err != nil {
				log.Printf("Error writing to file: %v", err)
			}
		}
	}

	duration := time.Since(start)
	workerJobStatistics.WorkerExecutionTime = duration.String()
	workerJobStatistics.TotalRequestsMade = totalRequestsMade

	fmt.Printf("Worker ID: %d processou no total: %d requisições\n", id, totalRequestsMade)

	err := writeFile(fmt.Sprintf("./responses/responses-from-worker-id-%d.json", id), workerJobStatistics)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
	}
}

func writeFile(filename string, data interface{}) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	numWorkers := runtime.NumCPU()
	var wg sync.WaitGroup

	fmt.Printf("Master cluster setting up %d workers...\n", numWorkers)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
