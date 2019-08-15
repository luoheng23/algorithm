package chapter2

import "testing"

func TestSelectSort(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	SelectSort(s[:])
	sortedS := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if s != sortedS {
		t.Errorf("SelectSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func TestSelectSortDecr(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	SelectSortDecr(s[:])
	sortedS := [10]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	if s != sortedS {
		t.Errorf("SelectSort failed. Expected %v, Got %v", sortedS, s)
	}
}

func BenchmarkSelectSort(b *testing.B) {
	const LENGTH = 10000
	var s [LENGTH]int
	for i := LENGTH; i > 0; i-- {
		s[LENGTH-i] = i
	}
	for i := 0; i < b.N; i++ {
		ss := s
		SelectSort(ss[:])
	}
}
