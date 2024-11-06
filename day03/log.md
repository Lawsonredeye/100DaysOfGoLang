# Understanding Functions and Return Values in GoLang - Day 03

Functions in GoLang are power tools in go that serves as a means of creating reuseable component which can be called even from different modules.

In go, you need to declare a reuseable component/function that can be access from another module it must be named with a Capital casing e.g ReuseableCode(), using lower casing makes the function accessible to the file it was created.

## How to create a function in GoLang
functions are created using the **func** keyword followed by a "()" which can take parameters, then followed by *"{}"* which is used to show the scope of the function.

Here a simple demo of a function in go:
```Go
func FuncThatReturnValue(age int) string {
	if age < 18 {
		return fmt.Sprintf("You are not qualified to drive at %d age.", age)
	}
	return "You are eligible to drive"
}
```

The function above takes an argument and returns a value (string value) when called.

## Multiple Parameter Function
A function can receive multiple parameters and it can be done by creating a function to receive more than one. Here is how to structure such:

```GO
func UserDetail(name string, age int, color string, balance float64) {
    ...
    // add logic here
}
```
As you can see, this function takes multiple parameters and when called the  total number of arguments must also be passed else the program would panic.

Also you can return multiple value from a function and here is how to create such:
```go
func UserDetail(name string, age int, color string, balance float64) (string, int, string, float64) {
    ...
    // add logic here
}
```

## Recursive Function
A function in GO can be recursive which is widely used in computer science to repeat a set of action till a condition is met.

Here is a simple recursive function which handles showcase that:

```Go
package main

import "fmt"

func recursion(num int) int {
    // this conditon makes the recursion to stop else the func would go on forever
    if num < 1 {
        return 1
    }
    fmt.Println(num)
    return recursion(num - 1)
}

func main() {
    // calling the func with an argument
    recursion(8)
}
```

## Named Function
This type of function has a return value in which its variable is declared and when just the *return* is called it returns the variable declared in the expected return value. Here is a demo:

```Go
func MyAge(age int) (result string) {
	result = fmt.Sprintf("lawson is %d years old", age)
	return
}
```
In the example above, result was declared in the return value section and when return was called the function automatically return the `result` variable even without being added to the return statement.

## Anonymous Function
Go has anonymous function which does not need to be named when created but the catch is that after being created it must be called immediately and it is not reuseable but best for cases where you need a quick functionality that wouldnt be needed anywhere else in the code.

Here is a demo:
```Go
package main

import "fmt"

func main()
func (){
    fmt.Println("Howdy, Neighbour")
    }() // passing the () after the closing bracket executes the function immediately.
```

## Reflection
I learnt more about how functions are being created and used in GoLang. The anonymous function seems a little bit off but i do know that further on we would find a good use for such function.

## Resources
1. [How to write create a function](https://www.w3schools.com/go/go_function_returns.php)
2. [Creating Functions (Net Ninja)](https://youtu.be/X68JmClzap4?si=hYEnVrSsZRmq250H)

