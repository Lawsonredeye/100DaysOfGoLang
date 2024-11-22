# Testing Testing Testing - Day 19

So far, i have been dwelling upon writing codes without testing them, which is a little bit ok since i am learning but i had to begin my journey on writing tests, even if it is just basic test cases.

I'll be delving into a topic known as Table Driven test which just uses a list of struct (inputs) and then run a loop to test all elements in the struct.

## Table Driven Test
Table Driven test in Go, are just regular Go functions (with a few rules). So what are the rules?

### TDT Rules
1. It needs to be in a file with a name like `xxx_test.go`

2. The test function must start with the word Test

3. The test function takes one argument only `t *testing.T`

4. To use the `*testing.T type`, you need to `import "testing"`, like you do with fmt in every file.

We'll write a unit test for this function starting with a file in the same directory, with the same package name, strings.
```Go
package demo

func Perimeter(width, height float64) float64 {
	return 2.0 * (width + height)
}
```

Below is the table driven test for the function defined above:
```Go
package demo

import (
    "reflect"
    "testing"
)

// using the test driven method to test the perim function
var perimtest = []struct {
	width    float64
	height   float64
	expected float64
}{
	{10.0, 10.0, 40.0},
	{20.0, 20.0, 80.0},
	{5.0, 5.0, 20.0},
	{25.1, 11.8, 73.8},
}

func TestPerimeter(t *testing.T) {
    // traversing through the slice to get the struct value
	for _, val := range perimtest {
		got := Perimeter(val.width, val.height)
		want := val.expected
    
    // using reflect to compare if both are truly equal
		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	}
}
```
To then test the function you run the command `go test` for golang to check for any `XXX_test.go` file and `go test -v` for a detailed output for the tests found.
This is just a basic touch up on testing with golang.


## Resources
1. [Table driven test](https://yourbasic.org/golang/table-driven-unit-test/)
2. [TDT - Dave Cheney](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
3. [Learn Go with tests - Recommended](https://quii.gitbook.io/learn-go-with-tests)

