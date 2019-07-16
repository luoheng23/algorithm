package chapter16

import "testing"

func TestGreedyActivitySelector(t *testing.T) {
	s := [11]int{1, 3, 0, 5, 3, 5, 6, 8, 8, 2, 12}
	f := [11]int{4, 5, 6, 7, 9, 9, 10, 11, 12, 14, 16}
	A := GreedyActivitySelector(s[:], f[:])
	if len(A) != 4 || A[0] != 0 || A[1] != 3 || A[2] != 7 || A[3] != 10 {
		t.Errorf("GreedyActivitySelector failed. Expected (0, 3, 7, 10), Got (%v)", A)
	}
}
