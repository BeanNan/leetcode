package utils

import "testing"

func TestEqualSliceInt1(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{1, 2}

	actual := EqualSliceInt(x, y)

	if actual != false {
		t.Errorf("slice equal error: %v %v", x, y)
	}
}

func TestEqualSliceInt2(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{1, 2, 4}

	actual := EqualSliceInt(x, y)

	if actual != false {
		t.Errorf("slice equal error: %v %v", x, y)
	}
}

func TestEqualSliceInt3(t *testing.T) {
	x := []int{1, 2, 3}
	y := []int{1, 2, 3}

	actual := EqualSliceInt(x, y)

	if actual != true {
		t.Errorf("slice equal error: %v %v", x, y)
	}
}
