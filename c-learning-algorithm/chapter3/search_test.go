package chapter3

import "testing"

func TestSearch(t *testing.T) {
	A := []int{9,10, 4, 4,23 ,2, 7, 21, 72}
	value1, value2 := 9, 20
	if Search(A, value1) != 0 || Search(A, value2) != -1 {
		t.Errorf("Search failed. Expect (%d, %d), Got (%d, %d).", 0, -1, Search(A, value1), Search(A, value2))
	}
}

func BenchmarkSearch(b *testing.B) {
	const n = 10000
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i
	}
	for i:= 0; i< b.N; i++ {
		Search(A, n)
	}
}
