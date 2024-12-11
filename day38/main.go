package main

import (
	"fmt"
	"sync"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	var m sync.Mutex

	go func() {
		m.Lock()
		defer m.Unlock()
		nIsEven := isEven(n)

		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, " is odd")
	}()

	go func() {
		m.Lock()
		defer m.Unlock()
		n++
	}()
	time.Sleep(time.Second)
}
