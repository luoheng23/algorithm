package chapter9

import "testing"

func TestRandSelect(t *testing.T) {
	s := [10]int{9, 7, 4, 3, 2, 6, 5, 8, 0, 1}
	m := RandSelect(s[:], 0, len(s), 5)
	if m != 4 {
		t.Errorf("RandSelect failed. Expected 4, Got %d", m)
	}
}
