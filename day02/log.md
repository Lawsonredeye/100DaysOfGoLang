# Understanding Defer, Panic and Effective Error handling - Day 02

Today i'll be learning how errors are being handled in GoLang to prevent code from either crashing during runtime due to variable.

## Errors
Errors can be caused by so many factors(human included) and also during compile time. It can be handled in go using the inbuilt **errors** library in golang.

Here's a simple way to handle an argument passed into a function.
```Go
package main
import "errors"

// sums the arg by itself or return an error when the arg is less than 0
func handleSumError(num int) (int, error) {
    if num < 0 {
        return 0, errors.New("number can't be less than 0")
    }
    return num + num, nil
}

func main() {
    // Calling the func and since handleSumError func returns an error and an int
    // the best way to handle both return values is to use two variables
    //to store the values to be returned by the function
    sum, err := handleSumError(5)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println(sum)
    }
}
```
Using the errors.New from the errors library, you can effectively return an error string when the handleSumError function is called if an argument doesnt satisfy the function logic.
One thing to note is that an error is either nil or a errors.New() string when being written in a function as that what the **error** expects.

## Defer
The Defer statement is a powerful statement as it can be used to perform clean-ups (for closing a file or closing the connection to a database) and also for handling concurrent actions when working with WaitGroups in goroutines.

For short, the defer statement delay the action of a function till the return or the end of a function where it is called is reached.

The defer keyword pushes an operation to a list of saved calls where it would be executed afer all operation at the stack is called off.

This is mostly useful in case an error occurs but you need to close a file during a file I/O operation.

Here's a code snippet on a function which utilizes the defer statement:
```Go
package main

import "fmt"

func main() {
    defer fmt.Println("This would print at the end")
    fmt.Println("This would printout first")
}
```

## Panic
The panic statement is used to stop the execution of a program and it is similar to a compile error.

It simply stops the ordinary flow of control and return an error to the console.

Here is a simple program that panics when the function is reached
```Go
func main() {
	fmt.Println("here comes the panic")
	panic("Error 204")
	fmt.Println("This wont run")
}
```
The program would stop executing when the panic() func is encountered and any code / logic after it wouldnt be reach.

It is just like calling return only that youre stopping the entire program from running.

## Recover
Can the Panic be stopped, well yes. The recover() func is used to recover a panicking program.

The recover() returns a nil if a panic() is found during the program and its used to revive it from panicking while the rest of the program continues which is being stored within the defered recover func() would be executed.

Here is how to use it:
```Go
package main

import "fmt"

func panicStopper() {
	if r := recover(); r != nil {
		fmt.Println("No panicking in this joint")
	}
	fmt.Println("I resumed here but not main")
	return
}

func main() {
	defer panicStopper()
	fmt.Println("here comes the panic")

	panic("Error 204")

    // This wouldnt run as panic has terminated the control flow of main. 
	fmt.Println("continuation")
}
```

By deferring panicStopper(), the func is moved to the saved call list and as soon as main is terminated (with panic of course), then the panicStopper would then be executed.


## Reflection
I learnt about how the defer, panic, error handling and recover() can be properly utilized when it comes to writing a program in Go.

I now understand that in I/O operations and concurrency the defer statement is extremely powerful for handling critical operations to prevent unexpected issues.
 
