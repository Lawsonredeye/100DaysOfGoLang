package main

import (
	"fmt"
)

func isEven(n int) bool {
	return n%2 == 0
}

// func main() {
// 	n := 0
// 	var m sync.Mutex

// 	go func() {
// 		m.Lock()
// 		defer m.Unlock()
// 		nIsEven := isEven(n)

// 		time.Sleep(5 * time.Millisecond)
// 		if nIsEven {
// 			fmt.Println(n, " is even")
// 			return
// 		}
// 		fmt.Println(n, " is odd")
// 	}()

// 	go func() {
// 		m.Lock()
// 		defer m.Unlock()
// 		n++
// 	}()
// 	time.Sleep(time.Second)
// }

func main() {
	bufferChan := make(chan int, 10)

	for i := 0; i < 10; i++ {
		go func() {
			bufferChan <- i
		}()
	}
	for i := 0; i < cap(bufferChan); i++ {
		select {
		case one := <-bufferChan:
			fmt.Println(one)
		case two := <-bufferChan:
			fmt.Println(two)
		}
	}
}
