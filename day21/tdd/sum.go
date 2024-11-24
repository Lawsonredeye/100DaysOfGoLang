package main

func Sum(num [5]int) int {
	sum := 0

	for _, val := range num {
		sum += val
	}
	return sum
}
