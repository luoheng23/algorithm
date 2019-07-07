package chapter2

import "testing"

func TestInsertionSort(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	InsertionSort(s[:])
	sortedS := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if s != sortedS {
		t.Errorf("InsertionSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func TestInsertionSortDecr(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	InsertionSortDecr(s[:])
	sortedS := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if s != sortedS {
		t.Errorf("InsertionSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	var s [100]int
	for i := 100; i > 0; i-- {
		s[100-i] = i
	}
	for i := 0; i < b.N; i++ {
		InsertionSort(s[:])
	}
}
