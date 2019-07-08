package chapter8

import "testing"

func TestCountSort(t *testing.T) {
	s := [20]int{0, 2, 4, 1, 0, 2, 5, 9, 8, 7, 5, 1, 2, 4, 3, 5, 9, 8, 7, 9}
	sortedS := [20]int{0, 0, 1, 1, 2, 2, 2, 3, 4, 4, 5, 5, 5, 7, 7, 8, 8, 9, 9, 9}
	CountSort(s[:], 9)
	if s != sortedS {
		t.Errorf("CountSort failed. Expected %v, Got %v", sortedS, s)
	}
}
