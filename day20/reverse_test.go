package main

import (
	"testing"
)

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
