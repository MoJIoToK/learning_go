package task33104

import (
	"math"
	"testing"
)

func TestSqrt(t *testing.T) {
	want := 3.0
	got := math.Sqrt(9)
	if got != want {
		t.Errorf("получено %f, ожидалось %f\n", got, want)
	}
}
