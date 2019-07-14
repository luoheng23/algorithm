package chapter2

func merge(A []int, p, q, r int) {
	const INTMAX = int(^uint(0) >> 1)
	A1, A2 := make([]int, q-p), make([]int, r-q)
	copy(A1, A[p:q])
	copy(A2, A[q:r])
	A1, A2 = append(A1, INTMAX), append(A2, INTMAX)
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

// MergeSort for array
func MergeSort(A []int, p, r int) {
	if r > p+1 {
		q := (r + p) / 2
		MergeSort(A, p, q)
		MergeSort(A, q, r)
		merge(A, p, q, r)
	}
}

func mergeSortGo(A []int, p, r int, m chan<- int) {
	if r > p+1 {
		q := (r + p) / 2
		if r-p > 1000 {
			a, b := make(chan int), make(chan int)
			go mergeSortGo(A, p, q, a)
			go mergeSortGo(A, q, r, b)
			<-a
			<-b
		} else {
			mergeSortGo(A, p, q, nil)
			mergeSortGo(A, q, r, nil)
		}
		merge(A, p, q, r)
		if m != nil {
			m <- 0
		}
	}
}

// MergeSortGo for goroutine array
func MergeSortGo(A []int, p, r int) {
	if r > p+1 {
		q := (r + p) / 2
		if r-p > 1000 {
			a, b := make(chan int), make(chan int)
			go mergeSortGo(A, p, q, a)
			go mergeSortGo(A, q, r, b)
			<-a
			<-b
		} else {
			MergeSortGo(A, p, q)
			MergeSortGo(A, q, r)
		}
		merge(A, p, q, r)
	}
}
