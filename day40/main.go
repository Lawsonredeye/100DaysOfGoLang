package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// create a simple program which handles a tasks/job such as a paymentService for a given online service using worker pools

// user choose service he wants to pay for and then click process
// the payment process begins and then the process

func ProcessServicePayment(jobs chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Processing your payment for %v\n", job)
		time.Sleep(3 * time.Second)
		fmt.Printf("Payment for %v has been completed!\n", job)
	}
}

var Services = [3]string{"Supersports", "Netflix", "Playstation Premium"}

func main() {
	jobs := make(chan string, 10)
	wg := sync.WaitGroup{}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go ProcessServicePayment(jobs, &wg)
	}

	for i := 1; i <= 9; i++ {
		jobs <- Services[rand.Intn(3)]
	}
	close(jobs)

	wg.Wait()
	fmt.Println("All payments have been processed successfully.")
}
