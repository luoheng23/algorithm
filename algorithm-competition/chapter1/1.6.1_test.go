package chapter1

import "testing"

func TestSolve(t *testing.T) {
	n, a := 5, []int{2, 3, 4, 5, 10}
	if solve(n, a) != 12 {
		t.Errorf("Solve failed. Expected 12, Got %d", solve(n, a))
	}
}
