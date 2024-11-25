# Pointers, Errors && Test Case

Pointers are very necessary when working in golang and you don't want to pass an argument by value but by reference to either a function or when creating methods.

It is one of the most powerful feature when working in golang as it is widely used with structs, functions, interfaces and even methods. 

We have touched this concept before but this time we will be adding test cases to our code while working with pointers.

Here is a basic way of creating a pointer in go:
```Go
package main

func main() {
    someNumber := 20
    ptr := &someNumber

    // checking the address of the pointer
    fmt.Printf("address of someNumber: %p, address of ptr: %p\n", &someNumber, ptr)
    // Output:
    // address of someNumber: 0xc000012070, address of ptr: 0xc00004c040 
}
```
By calling the `&` we were able to retrieve the memory address number of the variable someNumber. you could also do the same for the pointer and you would get its own different address.

## Quiz for the Day
1. Find the sum of integers from 1 to 10, from 20 to 37, and from 35 to 49, respectively, then total all three sum's.

## Resources
1. [Pointers, Errors && test](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors)

