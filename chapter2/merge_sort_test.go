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
	const LENGTH = 100000
	var s [LENGTH]int
	for i := LENGTH; i > 0; i-- {
		s[LENGTH-i] = i
	}
	for i := 0; i < b.N; i++ {
		ss := s
		MergeSort(ss[:], 0, len(s))
	}
}

func TestMergeSortGo(t *testing.T) {
	const LENGTH = 10000
	var s [LENGTH]int
	var sortedS [LENGTH]int
	for i := LENGTH; i > 0; i-- {
		s[LENGTH-i] = i
		sortedS[i-1] = i
	}
	MergeSortGo(s[:], 0, len(s))
	if s != sortedS {
		t.Errorf("MergeSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func BenchmarkMergeSortGo(b *testing.B) {
	const LENGTH = 100000
	var s [LENGTH]int
	for i := LENGTH; i > 0; i-- {
		s[LENGTH-i] = i
	}
	for i := 0; i < b.N; i++ {
		ss := s
		MergeSortGo(ss[:], 0, len(s))
	}
}
