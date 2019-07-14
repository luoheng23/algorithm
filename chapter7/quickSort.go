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

// QuickSort O(nlgn)
func QuickSort(A []int, p, r int) {
	if p+1 < r {
		q := partition(A, p, r)
		QuickSort(A, p, q)
		QuickSort(A, q+1, r)
	}
}

// RandQuickSort O(nlgn)
func RandQuickSort(A []int, p, r int) {
	if p+1 < r {
		q := randPartition(A, p, r)
		RandQuickSort(A, p, q)
		RandQuickSort(A, q+1, r)
	}
}

// RandQuickSortGo O(nlgn)
func RandQuickSortGo(A []int, p, r int, chans chan<- int) {
	if p+1 < r {
		q := randPartition(A, p, r)
		if r-p > 1000 {
			c := make(chan int, 2)
			go RandQuickSortGo(A, p, q, c)
			go RandQuickSortGo(A, q+1, r, c)
			for i := 0; i < 2; i++ {
				<-c
			}
		} else {
			RandQuickSortGo(A, p, q, nil)
			RandQuickSortGo(A, q+1, r, nil)
		}
	}
	if chans != nil {
		chans <- 0
	}
}

func partitionSameElement(A []int, p, r int) (int, int) {
	x, i, j := A[r-1], p, r
	for k := p; k < j; k++ {
		if A[k] <= x {
			A[i], A[k] = A[k], A[i]
			if A[i] != x {
				i++
			}
		} else if A[k] > x {
			j--
			A[j], A[k] = A[k], A[j]
		}
	}
	return i, j
}

func randPartitionSameElement(A []int, p, r int) (int, int) {
	i := rand.Intn(r-p) + p
	A[i], A[r-1] = A[r-1], A[i]
	return partitionSameElement(A, p, r)
}

// RandQuickSortSameElement O(nlgn)
func RandQuickSortSameElement(A []int, p, r int) {
	if p+1 < r {
		q, t := randPartitionSameElement(A, p, r)
		RandQuickSortSameElement(A, p, q)
		RandQuickSortSameElement(A, t, r)
	}
}

// TailRescursiveQuickSort O(nlgn)
func TailRescursiveQuickSort(A []int, p, r int) {
	for p+1 < r {
		q, t := randPartitionSameElement(A, p, r)
		TailRescursiveQuickSort(A, p, q)
		p = t
	}
}