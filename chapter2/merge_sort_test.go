package chapter2

import "testing"

func TestMergeSort(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	MergeSort(s[:], 0, len(s))
	sortedS := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if s != sortedS {
		t.Errorf("MergeSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	var s [100]int
	for i := 100; i > 0; i-- {
		s[100-i] = i
	}
	for i := 0; i < b.N; i++ {
		MergeSort(s[:], 0, len(s))
	}
}
