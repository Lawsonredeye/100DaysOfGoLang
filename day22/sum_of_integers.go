package main

func SumOfInteger() int {
	sum := 0

	for i := 1; i < 10; i++ {
		sum += i
	}

	for i := 20; i < 37; i++ {
		sum += i
	}

	for i := 37; i < 49; i++ {
		sum += i
	}
	return sum
}
