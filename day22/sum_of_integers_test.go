package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Golang quiz: Find the sum of integers from 1 to 10,
// from 20 to 37, and from 35 to 49, respectively,
// then total all three sum's.

// Solution:
// if respectively, when the loop counter gets to 10
// change the count to 20 && when the counter gets to 37
// change the count to 35 and the total length should be 49

func TestSumOfIntegers(t *testing.T) {
	got := SumOfInteger()
	want := 1031

	assert.Equal(t, want, got, "Total sum is not correct")
}
