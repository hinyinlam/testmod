package calculator

import "testing"

func TestMath(t *testing.T) {
	x := 1
	y := 2
	want := 3
	result := Add2(x, y)

	if result != want {
		t.Fatalf("x: %d, y: %d total should be: %d", x, y, want)
	}
}

func TestMultiple(t *testing.T) {
	var x = 10
	var y = 20
	var expect int64 = 200

	got := Multiple2(10, 20)
	if got != expect {
		t.Fatalf("X: %d, Y: %d, Expect: %d, Got: %d", x, y, expect, got)
	}

}
