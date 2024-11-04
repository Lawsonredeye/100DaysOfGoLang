package main

import "fmt"

func main() {
	// Printing to console..
	fmt.Println("Hello, Gophers!")

	// How to declare && initialize a variable
	var number int
	var balance float64
	var name string
	var isTrue bool
	var bytes []byte
	
	fmt.Println("default values of unintialized variable")
	fmt.Println(number) // default is 0
	fmt.Println(balance) // default is 0
	fmt.Println(name) // default is 0
	fmt.Println(isTrue) // default is false
	fmt.Println(bytes) // default is 0

	// adding values / initializing
	number = 12
	balance = 828.9
	name = "lawson"
	isTrue = true
	bytes = []byte(name) 

	fmt.Println("variables with values")
	fmt.Println(number)
	fmt.Println(balance)
	fmt.Println(name) 
	fmt.Println(isTrue) 
	fmt.Println(bytes) 
	
	if number > 12 {
		fmt.Println("number is too high")
	} else if number < 12 {
		fmt.Println("number is too low")
	} else {
		fmt.Println("number is just right")
	}
}
