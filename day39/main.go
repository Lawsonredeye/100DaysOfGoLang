package main

import (
	"fmt"
	"time"
)

// Worker processor for handling the job passed into the channel
func worker(id int, jobs <-chan int, results chan<- int) {

	// loops through the jobs channel and then carry out a task.
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// Simulates the number of jobs to be carried out by the goroutine
	jobs := make(chan int, 100)

	// results from each jobs
	results := make(chan int, 100)

	// Spawns 3 workers for handling the jobs
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// fills the jobs channel with the the value of j
	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
}
