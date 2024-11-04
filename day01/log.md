# Learning The Basics of Go - Day 01

GoLang is a very simple language to grasp even if you've never learned about programming before.

It's syntax is similar to that of C && Python when you look really close to it.

It is a statically typed language but can still give you a feel of how dynamic typed language (the likes of Python, Javascript).

## Print to Console
There are several ways to print to the console but for today the basics would be to use one of the standard library tool (std tool)  the "fmt" package to print to the console.

Here is how:
```Go
package main
import "fmt" // functions or libraries to be used must be imported using the import keyword and enclosed in double quotes

func main() {
    fmt.Println("Hello, Gophers!")
}
```

**Note:**

All program written here should be placed inside a  `func main(){}` function as the main is the point of execution in GoLang

## How to declare a variable in GoLang

there are two ways of declaring a variable for any values you need to create:

```Go
// var [name of variable] [data type of variable]
var number int
```
and you can declare using the shorthand rule where you not only declare but initialize the variable all together.
Here's the syntax:
```Go
number := 56
name := "lawson" // string based values are always in double quotes
```

## Data types
Data types are just the type of values that a variable in go can hold and store.
And since Go is statically typed (meaning you have to declare the type of data you wanna store) you cant change the type of data after you have created it as it must be known before compile time.

Here are the major data types in go:
```Go
var booleans bool // stores true or false
var stringsValue string // stores string
var integers int //stores integers in 64 bits (by default) others are int8, int16, int32, int64
var floatPoint float64 // stores floatpoint numbers eg decimal point 10.3, 0.2, 5.27..  others are float32
var bytesValue []bytes // stored data in bytes
var pointer uintptr // stores address of a variable
```

And here is how to initialize them:
boolean := true // value can be in true or false
number := 32 // by default its int64
balance := 2415.0 // float64 by default
name := "lawsonredeye" // must be enclosed in double quotes

## Conditionals

Just like every other language, go has its own way of writing conditionals for handling specific cases as well as in terms of error handling.

It uses the traditional `if..else, if..else if && if..else if .. else` , construct

Here's a simple way of writing a simple condition based on the age of a person who can obtain a drivers license.

```Go
studentAge := 17

if studentAge < 18 {
    fmt.Println("Student is unable to obtain a license")
} else if studentAge > 18 && studentAge < 75 {
    fmt.Println("Student is able to obtain a license")
} else {
    fmt.Println("Student is unable to drive and obtain a license")
}
```
One thing to note is that Go doesnt have a tenary operator like JavaScript or Python or even like C/C++.

## Reflection
I understood some basic concept of declaring and intializing a variable in golang, how to use conditionals in golang, how to print to console (stdout) and also some other data types.

Today, I learned the basics of Go syntax and how to declare variables. I found the distinction between static and dynamic typing particularly interesting and am excited to explore more!

### Links
[learn about go bytes](https://golangdocs.com/golang-bytes-variable-package)
[All data types in GoLang](https://golangbyexample.com/all-data-types-in-golang-with-examples/)

