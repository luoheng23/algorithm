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

func TestMaxHeap(t *testing.T) {
	s := [9]int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	sortedS := [9]int{25, 20, 17, 13, 8, 7, 5, 4, 2}
	newS := [9]int{}
	h := BuildMaxHeap(s[:])
	for i := 0; i < 9; i++ {
		newS[i], _ = h.ExtractMax()
	}
	if newS != sortedS {
		t.Errorf("MaxHeap failed. Expected %v, Got %v.", sortedS, newS)
	}

}

func TestMinHeap(t *testing.T) {
	s := [9]int{5, 13, 2, 25, 7, 17, 20, 8, 4}
	sortedS := [9]int{2, 4, 5, 7, 8, 13, 17, 20, 25}
	newS := [9]int{}
	h := BuildMinHeap(s[:])
	for i := 0; i < 9; i++ {
		newS[i], _ = h.ExtractMin()
	}
	if newS != sortedS {
		t.Errorf("MaxHeap failed. Expected %v, Got %v.", sortedS, newS)
	}

}
