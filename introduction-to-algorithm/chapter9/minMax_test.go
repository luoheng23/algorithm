package chapter9

import "testing"

func TestMin(t *testing.T) {
	s := [10]int{5, 9, 3, 4, 2, 1, 0, 7, 8, 6}
	m := Min(s[:])
	if m != 0 {
		t.Errorf("Min failed. Expected 0, Got %d", m)
	}
}

func TestMax(t *testing.T) {
	s := [10]int{5, 9, 3, 4, 2, 1, 0, 7, 8, 6}
	m := Max(s[:])
	if m != 9 {
		t.Errorf("Min failed. Expected 9, Got %d", m)
	}
}

func TestMinMax(t *testing.T) {
	s := [10]int{5, 9, 3, 4, 2, 1, 0, 7, 8, 6}
	min, max := MinMax(s[:])
	if min != 0 || max != 9 {
		t.Errorf("MinMax failed. Expected (0, 9), Got (%d, %d)", min, max)
	}
}
