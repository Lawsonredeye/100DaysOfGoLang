package main

import "fmt"

func sendOnlyChan(ch chan<- int) {
	ch <- 10
}

func main() {
	ch := make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch <- i
	}
	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(cap(ch))
	fmt.Println(<-ch)

	// unidirectional channels that can only be written to and not read from
	uniChan := make(chan<- int)
	go sendOnlyChan(uniChan)

	//fmt.Println(<-uniChan) // Would panic as uniChan can only be written to and not read from

	biDirectionalChan := make(chan int)

	go sendOnlyChan(biDirectionalChan)
	fmt.Println(<-biDirectionalChan) // Works as the channel can be written to and read from
}
