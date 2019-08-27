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

func TestBinarySearch(t *testing.T) {
	A := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	value1, value2 := 9, 20
	if BinarySearch(A, value1) != 9 || BinarySearch(A, value2) != -1 {
		t.Errorf("BinarySearch failed. Expect (%d, %d), Got (%d, %d).", 9, -1, BinarySearch(A, value1), BinarySearch(A, value2))
	}
}

func BenchmarkBinarySearch(b *testing.B) {
	const n = 10000
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = i
	}
	for i:= 0; i< b.N; i++ {
		BinarySearch(A, n)
	}
}
