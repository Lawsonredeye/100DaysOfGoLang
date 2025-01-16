package main

import (
	"fmt"
)

func main() {
	fmt.Println(dayOfProgrammer(2000))
}

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {

	// if the calender was changed at 1919
	// subtract 28 - 14 => 14
	// subtract 244 - 14 => 230
	// subtract 256 - 230 => 26
	// therefore 26.09.1918 was our day
	if year == 1918 {
		return fmt.Sprintf("26.09.%d", year)
	}

	leapYear := isLeapYear(year)

	if leapYear {
		return fmt.Sprintf("12.09.%d", year)
	}
	return fmt.Sprintf("13.09.%d", year)
}

func isLeapYear(year int32) bool {

	if year <= 1919 {
		if year%4 == 0 {
			return true
		}
	}

	if year%4 == 0 && year%100 != 0 {
		return true
	}

	if year%400 == 0 && year%100 == 0 {
		return true
	}

	return false
}
