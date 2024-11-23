package main

import "fmt"

// Reverse the string and remove numbers from the range of 0-9

// func ModifyString(s string) string {
// 	// loop through the string to see if there are numbers in it and then remove the numbers from the string.

// 	for idx := 0; idx < 10; idx++ {
// 		for ldx := idx + 1; ldx < len(); ldx++ {
// 			if
// 		}
// 	}
// }

func main() {
	str := "noswal"
	reversed := Reversed(str)
	fmt.Println(reversed)
}

func Reversed(s string) string {
	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return reversed
}
