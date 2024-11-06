package main

import "fmt"

// VoidFunc is importable because it begins with an Upper Case
func VoidFunc() {
	fmt.Println("This takes no value but prints out a string")
}

func FuncWithArgs(name string) {
	fmt.Printf("This takes an argument and prints it to the console")
}

func FuncWithReturnValue(age int) string {
	if age < 18 {
		return fmt.Sprintf("You are not qualified to drive at %d age.", age)
	}
	return "You are eligible to drive"
}

func RecursiveFunction(num int) int {
	if num == 20 {
		return 0
	}
	fmt.Println(num)
	return RecursiveFunction(num + 1)
}

func main() {
	// func that takes no argument but prints to the console
	VoidFunc()

	// func receives an argument but returns no value
	FuncWithArgs("Lawson")

	// Function that returns a value must be stored inside a variable
	res := FuncWithReturnValue(28)
	fmt.Println(res)

	// A simple recursive function
	RecursiveFunction(10)

	// Anonymous func
	func (){
		fmt.Println("Hi, I am an anonymous func")
	}()
}
