package main

import "fmt"


func deferPrintOperation(name string) {
	greeting := fmt.Sprintf("hello, %s! Welcome to golang", name)
		defer fmt.Println(greeting)
	return
}

func panicStopper() {
	if r := recover(); r != nil {
		fmt.Println("No panicking in this joint")
	}
	fmt.Println("I resumed here but not main")
	return
}

func main() {
	deferPrintOperation("Lawson redeye") // output: hello, Lawson redeye! Welcome to golang
	defer panicStopper()
	fmt.Println("here comes the panic")
	panic("Error from day 02")

	// This wouldnt run as panic has terminated the control flow of main. 
	fmt.Println("continuation")
	fmt.Println("Hello, main.goroutine!")
}

