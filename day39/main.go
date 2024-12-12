package main

import (
	"fmt"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("Successfully wrote", i, "to ch")
	}
	close(ch)
}

// Buffered Channels
// func main() {
// 	// ch := make(chan string, 2)
// 	// ch <- "lawson"
// 	// ch <- "redeye"

// 	// fmt.Println(<-ch)
// 	// fmt.Println(<-ch)
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)

// 	for v := range ch {
// 		fmt.Println("Read value", v, "from ch")
// 		time.Sleep(2 * time.Second)
// 	}
// }

// Creating deadlocks with channels
func main() {
	ch := make(chan string, 2)

	ch <- "lawson"
	ch <- "redeye"
	ch <- "Agent"
	fmt.Println(<-ch)
}
