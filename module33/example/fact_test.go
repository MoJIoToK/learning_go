package example

import "testing"

func TestFact(t *testing.T) {
	want := 6
	got := Fact(3)
	if got != want {
		t.Errorf("Fact() = %v, want %v", got, want)
	}
}
