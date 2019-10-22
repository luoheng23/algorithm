package chapter3

import "testing"

func TestMinI(t *testing.T) {
	a, K := []int{2, 3, 3, 5, 6}, 3
	if minI(a, K) != 1 {
		t.Errorf("TestMinI failed. Expected 1, Got %d.", minI(a, K))
	}
}
