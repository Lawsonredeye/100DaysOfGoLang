package main

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("expected: %v, got: %v, given: %v\n", want, got, numbers)
	}
}

func TestSumWithRun(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{2, 4, 6, 8, 10}

		got := Sum(numbers)
		want := 30

		if got != want {
			t.Errorf("expected: %v, got: %v, given: %v\n", want, got, numbers)
		}
	})

	t.Run("testing slice of array", func(t *testing.T) {
		sliceArr := []int{3, 6, 9, 12, 15}

		got := Sum(sliceArr)
		want := 45

		if got != want {
			t.Errorf("expected: %v, got: %v, given: %v\n", want, got, sliceArr)
		}
	})
}
