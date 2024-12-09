# WaitGroups

waitgroups blocks execution of a function until its internal counter becomes 0.

They are useful when working with goroutines as it helps to ensure that all goroutines are completely executed in a concurrent environment.

Without WaitGroups, goroutines would not get executed in a function.

## Using time.Sleep() for Pausing Main Goroutine
```GO
package main

import (
	"fmt"
	"time"
)

func main() {
	go helloworld()
	go goodbye()
	time.Sleep(2 * time.Second)
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}
```
**Simple go program which utilizes the time.Sleep() method for pausing execution**

Using the time.Sleep is good for a small scale program but when it comes to large scale programs its best to use sync.WaitGroups to handle concurrency.

## Using WaitGroup for Handling Concurrency
WaitGroups are perfect for handling concurrency without using the time.Sleep for pausing execution which makes it best for handling much concurrent requests.
Here is a simple program to demostrate how to use:
```Go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

    // Helps waitgroup to know the number of goroutines for it to handle
	wg.Add(2)
	go helloworld(&wg)
	go goodbye(&wg)

    // pauses all to ensure that the number of waitgroup counter has been exhausted
	wg.Wait()
}

func helloworld(wg *sync.WaitGroup) {
    // Reduce the waitgroup counter by 1
	defer wg.Done()
	fmt.Println("Hello World!")
}

func goodbye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Good Bye!")
}
```



