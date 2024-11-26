package math

import (
	"testing"
)

// Table Driven Test for the Add func() in our Math library.
func TestAdd(t *testing.T) {
	rangeNums := []struct {
		testName    string
		leftAddend  int
		rightAddend int
		want        int
	}{
		{"less than 10", 5, 8, 13},
		{"less than 20", 19, 12, 31},
		{"negative numbers", -56, 75, 19},
		{"both negative", -25, -25, -50},
	}

	for _, tt := range rangeNums {
		got := Add(tt.leftAddend, tt.rightAddend)
		if got != tt.want {
			t.Errorf("%v: got: %v, want: %v", tt.testName, got, tt.want)
		}
	}
}

func TestSubtract(t *testing.T) {
	got := Subtract(10, 2)
	want := 8
	if got != want {
		t.Errorf("basic subtract: got: %v, want: %v", got, want)
	}
}
