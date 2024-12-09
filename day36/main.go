package main

import (
	"fmt"
)

func main() {
	go helloworld()
	go goodbye()
}

func helloworld() {
	fmt.Println("Hello World!")
}

func goodbye() {
	fmt.Println("Good Bye!")
}
