package chapter2

func merge(A []int, p, q, r int) {
	const INT_MAX = int(^uint(0) >> 1)
	A1, A2 := make([]int, q-p), make([]int, r-q)
	copy(A1, A[p:q])
	copy(A2, A[q:r])
	A1, A2 = append(A1, INT_MAX), append(A2, INT_MAX)
	i, j, k := 0, 0, 0
	for k = p; k < r; k++ {
		if A1[i] >= A2[j] {
			A[k] = A2[j]
			j++
		} else {
			A[k] = A1[i]
			i++
		}
	}
}

func MergeSort(A []int, p, r int) {
	if r > p+1 {
		q := (r + p) / 2
		MergeSort(A, p, q)
		MergeSort(A, q, r)
		merge(A, p, q, r)
	}
}
