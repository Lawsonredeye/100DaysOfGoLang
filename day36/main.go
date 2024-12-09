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
