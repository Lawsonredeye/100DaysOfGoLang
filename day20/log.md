# Practicing Test Driven Test

So, yesterday i began writing test cases to strengthen my ability to write functions which returns the predicted values.

So today, i'll be delving into different ways of utilizing tests and seeing how far i can push my self in writing tests. Let's begin.

To write tests cases for your project, you need to always remember to create a file that ends with `xxx_test.go` as it would help the `go test` command to detect your test codes easily.

Also a side note, you must have a module for you to be able to run `go test` command.

For today we will create a func() that reverse a string and then write some basic test for it. Below is the function

```Go
// reverse.go

package demo

func Reversed(s string) string {
	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return reversed
}
```
The test case for the reverse func()
```Go
// reverse_test.go

package main

import "testing"

func TestReverse(t *testing.T) {
	testString := []struct {
		testName string
		str      string
		want     string
	}{
		{"reverse capitals", "YTIC", "CITY"},
		{"lower case", "airegin", "nigeria"},
		{"Sentence case", "notsoB", "Boston"},
	}

	for _, tt := range testString {
		got := Reversed(tt.str)

		if got != tt.want {
			t.Fatalf("%v: expected: %v, want: %v\n", tt.testName, got, tt.want)
		}
	}
}
```

we used the `table driven test` format for testing the `Reverse` function, where we created a `[]struct` with some few test cases and also added members which distinguish each test in the case of failure.

## Assignment
1. Write a basic test for a function that returns `hello world` to a variable
 2. Make sure to use table driven test for testing the function multiple times.
