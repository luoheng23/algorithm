package chapter1

import "testing"

func TestAnts(t *testing.T) {
	n, L, x := 3, 10, []int{2, 6, 7}
	min, max := ants(n, L, x)
	if min != 4 || max != 8 {
		t.Errorf("Ants failed. Expected (4, 8), Got (%d, %d).", min, max)
	}
}
