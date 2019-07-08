package chapter7

import "math/rand"

func partition(A []int, p, r int) int {
	x, i := A[r-1], p
	for j := p; j < r-1; j++ {
		if A[j] <= x {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}
	A[i], A[r-1] = A[r-1], A[i]
	return i
}

func randPartition(A []int, p, r int) int {
	i := rand.Intn(r-p) + p
	A[i], A[r-1] = A[r-1], A[i]
	return partition(A, p, r)
}

func QuickSort(A []int, p, r int) {
	if p+1 < r {
		q := partition(A, p, r)
		QuickSort(A, p, q)
		QuickSort(A, q+1, r)
	}
}

func RandQuickSort(A []int, p, r int) {
	if p+1 < r {
		q := randPartition(A, p, r)
		QuickSort(A, p, q)
		QuickSort(A, q+1, r)
	}
}
