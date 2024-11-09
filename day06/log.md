# Pointers - Day 06

Pointers are variables that points to a memory address and stores that address.

It is mainly useful in complex data structures like linkedlist and for  passing data by reference and not just values to a function.

Here is how to create a pointer in Go.
```Go
number := 32
ptr := &number // using & returns the pointer to a variable
```
In to access the value from that pointer in which the pointer has the address to, you use the dereference operator `(*)` to get the value and to set the value.

Pointers are very useful if youre working with a function and dont want to work with a copy.

You can choose to pass in the address of the argument into the function and it would manipulate the data in which the address was passed into it.
Here is a simple example of how pointers works with functions:
```Go
package main
func SumReference(num *int) {
    *num += *num 
}

func main() {
    number := 5
    SumReference(&number)
    fmt.Println(number)
    // Output: 10
}
```
You could see that number automatically becomes 10 even tho we didn't call the return statement on the function.
This is possible because we manipulated the data in that address of the argument that was passed to the function and not a copy of the argument.

Although pointers seems cool but accessing a nil pointer can be dangerous and could cause the compiler to panic.

## Resources
1. [Pointers](https://go.dev/tour/moretypes/1)
2. [video on Pointer](https://www.youtube.com/watch?v=a4HcEsJ1hIE)
