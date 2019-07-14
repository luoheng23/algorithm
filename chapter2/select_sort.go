package chapter2

// SelectSort for array
func SelectSort(A []int) {
	for j := 0; j < len(A); j++ {
		min := j
		for i := j+1; i < len(A); i++ {
			if A[min] > A[i] {
				min = i
			}
		}
		A[j], A[min]  = A[min], A[j]
	}
}

// SelectSortDecr for array
func SelectSortDecr(A []int) {
	for j := 0; j < len(A); j++ {
		max := j
		for i := j+1; i < len(A); i++ {
			if A[max] < A[i] {
				max = i
			}
		}
		A[j], A[max] = A[max], A[j]
	}
}
