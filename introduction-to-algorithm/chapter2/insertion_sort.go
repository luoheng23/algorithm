package chapter2

// InsertionSort O(n^2)
func InsertionSort(A []int) {
	for j := 1; j < len(A); j++ {
		key, i := A[j], j-1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}

// InsertionSortDecr O(n^2)
func InsertionSortDecr(A []int) {
	for j := 1; j < len(A); j++ {
		key, i := A[j], j-1
		for i >= 0 && A[i] < key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}
