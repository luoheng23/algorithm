package chapter8

// InsertionSortForBucketSort O(n^2)
func InsertionSortForBucketSort(A []float64) {
	for j := 1; j < len(A); j++ {
		key, i := A[j], j-1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}
}

// BucketSort O(n)
func BucketSort(A []float64) {
	n := len(A)
	B := make([][]float64, n)
	fN := float64(n)
	for _, a := range A {
		B[int(a * fN)] = append(B[int(a * fN)], a)
	}
	for _, b := range B {
		InsertionSortForBucketSort(b)
	}
	i := 0
	for j:= 0; j < n; j++ {
		m := len(B[j])
		for k := 0; k < m; k++ {
			A[i] = B[j][k]
			i++
		}
	}
}