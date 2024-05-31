package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// Task struct represents a unit of work.
type Task struct {
	ID      int
	Payload string
}

func add(a int, b int) int {

}

// Process simulates processing the task.
func (t Task) Process() (string, error) {
	if t.Payload == "" {
		return "", errors.New("empty payload")
	}
	time.Sleep(100 * time.Millisecond) // Simulate processing time
	return fmt.Sprintf("Processed task %d: %s", t.ID, t.Payload), nil
}

// WorkerPool struct to manage a pool of workers.
type WorkerPool struct {
	tasks   []Task
	results chan string
	errors  chan error
	wg      sync.WaitGroup
	sem     chan struct{}
}

// NewWorkerPool creates a new WorkerPool.
func NewWorkerPool(tasks []Task, maxWorkers int) *WorkerPool {
	return &WorkerPool{
		tasks:   tasks,
		results: make(chan string, len(tasks)),
		errors:  make(chan error, len(tasks)),
		sem:     make(chan struct{}, maxWorkers),
	}
}

// Run starts the worker pool.
func (wp *WorkerPool) Run() {
	for _, task := range wp.tasks {
		wp.wg.Add(1)
		wp.sem <- struct{}{} // Acquire a semaphore slot
		go wp.processTask(task)
	}

	wp.wg.Wait()
	close(wp.results)
	close(wp.errors)
}

// processTask processes a single task.
func (wp *WorkerPool) processTask(task Task) {
	defer wp.wg.Done()
	defer func() { <-wp.sem }() // Release semaphore slot

	result, err := task.Process()
	if err != nil {
		wp.errors <- err
	} else {
		wp.results <- result
	}
}

// Results returns the channel to read results from.
func (wp *WorkerPool) Results() <-chan string {
	return wp.results
}

// Errors returns the channel to read errors from.
func (wp *WorkerPool) Errors() <-chan error {
	return wp.errors
}

func Example1() {
	tasks := []Task{
		{ID: 1, Payload: "Task 1 data"},
		{ID: 2, Payload: "Task 2 data"},
		{ID: 3, Payload: ""},
		{ID: 4, Payload: "Task 4 data"},
		{ID: 5, Payload: "Task 5 data"},
	}

	workerPool := NewWorkerPool(tasks, 3)
	go workerPool.Run()

	for result := range workerPool.Results() {
		fmt.Println(result)
	}

	for err := range workerPool.Errors() {
		fmt.Println("Error:", err)
	}
}
