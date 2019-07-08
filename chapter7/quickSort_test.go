package chapter7

import "testing"

func TestQuickSort(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	QuickSort(s[:], 0, 10)
	sortedS := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if s != sortedS {
		t.Errorf("InsertionSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	const LENGTH = 10000
	var s [LENGTH]int
	for i := 0; i < LENGTH; i++ {
		s[LENGTH-i-1] = i
	}
	for i := 0; i < b.N; i++ {
		ss := s
		QuickSort(ss[:], 0, 10)
	}
}
func TestRandQuickSort(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	RandQuickSort(s[:], 0, 10)
	sortedS := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if s != sortedS {
		t.Errorf("InsertionSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func BenchmarkRandQuickSort(b *testing.B) {
	const LENGTH = 10000
	var s [LENGTH]int
	for i := 0; i < LENGTH; i++ {
		s[LENGTH-i-1] = i
	}
	for i := 0; i < b.N; i++ {
		ss := s
		RandQuickSort(ss[:], 0, 10)
	}
}

func TestRandQuickSortGo(t *testing.T) {
	const LENGTH = 10000
	var s [LENGTH]int
	var ss [LENGTH]int
	for i := 0; i < LENGTH; i++ {
		s[LENGTH-i-1] = i
		ss[i] = i
	}
	RandQuickSortGo(s[:], 0, LENGTH, nil)
	if s != ss {
		t.Errorf("InsertionSort failed. Expected %v, Got %v", ss, s)
	}
}

func BenchmarkRandQuickSortGo(b *testing.B) {
	const LENGTH = 10000
	var s [LENGTH]int
	for i := 0; i < LENGTH; i++ {
		s[LENGTH-i-1] = i
	}
	for i := 0; i < b.N; i++ {
		ss := s
		RandQuickSortGo(ss[:], 0, 10000, nil)
	}
}
