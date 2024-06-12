package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a unit of work to be processed by a worker.
type Job struct {
	ID    int
	Value int
}

// Result represents the outcome of processing a Job.
type Result struct {
	JobID    int
	Output   int
	WorkerID int
}

// Worker function processes jobs and sends results to the results channel.
func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Simulate processing time
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(10000)))
		output := job.Value * 2 // Simple processing: double the value
		results <- Result{JobID: job.ID, Output: output, WorkerID: id}
	}
}

// Dispatcher function creates workers and assigns jobs to them.
func dispatcher(numWorkers int, jobs <-chan Job, results chan<- Result) {
	var wg sync.WaitGroup
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	wg.Wait()

	close(results)
}

func PracticeChannel() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Create job and result channels
	numJobs := 10
	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	// Start the dispatcher with 3 workers
	go dispatcher(3, jobs, results)

	// Send jobs to the job channel
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, Value: rand.Intn(100)}
	}
	close(jobs)

	// Collect results from the results channel
	for result := range results {
		fmt.Printf("Job ID: %d, Value: %d, Processed by Worker ID: %d, Result: %d\n",
			result.JobID, result.JobID, result.WorkerID, result.Output)
	}
}
