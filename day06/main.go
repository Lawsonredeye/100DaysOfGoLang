package main

import "fmt"


func main() {
	// creating a pointer to a variable
	name := "lawson"
	var namePointer *string

	namePointer = &name

	*namePointer = "lawson redeye"

	fmt.Printf("Memory address of name: %p\n", &name)
	fmt.Printf("Memory address of namePointer: %p\n", namePointer)
	fmt.Printf("Data on name: %s, data on namePointer: %v", name, *namePointer)
}
