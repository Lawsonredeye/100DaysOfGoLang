# Testing Testing Testing - Day 19

So far, i have been dwelling upon writing codes without testing them, which is a little bit ok since i am learning but i had to begin my journey on writing tests, even if it is just basic test cases.

I'll be delving into a topic known as Table Driven test which just uses a list of struct (inputs) and then run a loop to test all elements in the struct.

## Table Driven Test
Table Driven test in Go, are just regular Go functions (with a few rules).

```Go
// Similar to the split func in the strings library.
func Split(s, sep string) []string {
	var result []string

	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	return append(result, s)
}
```

We'll write a unit test for this function starting with a file in the same directory, with the same package name, strings.



## Resources
1. [Table driven test](https://yourbasic.org/golang/table-driven-unit-test/)
2. [TDT - Dave Cheney](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)


