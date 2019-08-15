package chapter5

import "testing"

func TestShuffle(t *testing.T) {
	count := 0
	for i := 0; i < 10; i++ {
		var A = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		B := A
		Shuffle(A[:])
		if A == B {
			count++
		}
	}
	if count >= 2 {
		t.Errorf("Shuffle failed. Expected (0, 1, 2), Got %d", count)
	}
}
