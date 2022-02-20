package calculator

import "testing"

func TestMath(t *testing.T) {
	x := 1
	y := 2
	want := 3
	result := Add(x, y)

	if result != want {
		t.Fatalf("x: %d, y: %d total should be: %d", x, y, want)
	}
}
