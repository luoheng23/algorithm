package chapter6

import "testing"

func TestHeapSort(t *testing.T) {
	s := [9]int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	sortedS := [9]int{2, 4, 5, 7, 8, 13, 17, 20, 25}
	HeapSort(s[:])
	if s != sortedS {
		t.Errorf("HeapSort failed. Expected %v, Got %v.", sortedS, s)
	}
}
