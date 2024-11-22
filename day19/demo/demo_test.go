package split

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	got := Split("a/b/c", "/")
	want := []string{"a", "b", "c"}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}

// using the test driven method to test the perim function
var perimtest = []struct {
	width    float64
	height   float64
	expected float64
}{
	{10.0, 10.0, 40.0},
	{20.0, 20.0, 80.0},
	{5.0, 5.0, 20.0},
	// {25.1, 11.8, 73.80000000001},
}

func TestPerimeter(t *testing.T) {
	// got := Perimeter(10.0, 10.0)
	// want := 40.0
	for _, val := range perimtest {
		got := Perimeter(val.width, val.height)
		want := val.expected

		if !reflect.DeepEqual(want, got) {
			t.Fatalf("expected: %v, got: %v", want, got)
		}
	}
}

func TestArea(t *testing.T) {
	got := Area(4, 6)
	want := 12

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
