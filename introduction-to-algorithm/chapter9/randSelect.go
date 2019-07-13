package chapter9

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

// choose the ith element
func RandSelect(A []int, p, r, i int) int {
	if p+1 == r {
		return A[p]
	}
	q := randPartition(A, p, r)
	k := q - p + 1
	if i == k {
		return A[q]
	} else if i < k {
		return RandSelect(A, p, q, i)
	} else {
		return RandSelect(A, q+1, r, i-k)
	}

}
