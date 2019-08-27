package chapter3

import (
	"sort"
	"testing"
)

func TestShellSort(t *testing.T) {
	A := [...]int{20, 17, 4, 9, 3, 5, 22, 100, 88, 77, 99, 56, 42, 73}
	B := A
	sort.Ints(B[:])
	ShellSort(A[:], 5)
	if A != B {
		t.Errorf("ShellSort failed. Expected %v, Got %v", B, A)
	}
}

func BenchmarkShellSort(b *testing.B) {
	const n = 10000
	A := make([]int, n)
	for i := 0; i < n; i++ {
		A[i] = n - i
	}
	for i := 0; i < b.N; i++ {
		ShellSort(A, 5)
	}
}
