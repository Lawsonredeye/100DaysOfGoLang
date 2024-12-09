package main

import (
	"fmt"
	"time"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go helloworld(&wg)
// 	go goodbye(&wg)
// 	wg.Wait()
// }

// func helloworld(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Hello World!")
// }

// func goodbye(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println("Good Bye!")
// }

func main() {
	msg := make(chan string)
	go greet(msg)
	time.Sleep(2 * time.Second)

	greeting := <-msg

	time.Sleep(2 * time.Second)
	fmt.Println("Greeting recieved")
	fmt.Println(greeting)

	_, ok := <-msg
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!")
	}
}

func greet(ch chan string) {
	fmt.Println("Greeter waiting to send greetings!")
	ch <- "Hello Docker"
	close(ch)

	fmt.Println("Greeter completed")
}
