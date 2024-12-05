package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		GoroutineFunc()
	}()
	fmt.Println("Session ended..")
	wg.Wait()
}

func GoroutineFunc() {
	fmt.Println("Executed by a goroutine... Hehe")
}
