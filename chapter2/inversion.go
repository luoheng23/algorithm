package chapter2

func count(A []int, p, q, r int) int {
	const INTMAX = int(^uint(0) >> 1)
	A1, A2 := make([]int, q-p), make([]int, r-q)
	copy(A1, A[p:q])
	copy(A2, A[q:r])
	A1, A2 = append(A1, INTMAX), append(A2, INTMAX)
	i, j, k := 0, 0, 0
	count := 0
	n1 := len(A1) - 1
	for k = p; k < r; k++ {
		if A1[i] > A2[j] {
			A[k] = A2[j]
			j++
			count += n1 - i
		} else {
			A[k] = A1[i]
			i++
		}
	}
	return count
}

// MergeCount O(n^2)
func MergeCount(A []int, p, r int) int {
	c := 0
	if r > p+1 {
		q := (r + p) / 2
		c += MergeCount(A, p, q)
		c += MergeCount(A, q, r)
		c += count(A, p, q, r)
	}
	return c
}

func mergeCountGo(A []int, p, r int, m chan<- int) int {
	c := 0
	if r > p+1 {
		q := (r + p) / 2
		if r-p > 1000 {
			a, b := make(chan int), make(chan int)
			go mergeCountGo(A, p, q, a)
			go mergeCountGo(A, q, r, b)
			c += <-a
			c += <-b
		} else {
			c += mergeCountGo(A, p, q, nil)
			c += mergeCountGo(A, q, r, nil)
		}
		c += count(A, p, q, r)
		if m != nil {
			m <- c
		}
	}
	return c
}

// MergeCountGo O(n^2)
func MergeCountGo(A []int, p, r int) int {
	c := 0
	if r > p+1 {
		q := (r + p) / 2
		if r-p > 1000 {
			a, b := make(chan int), make(chan int)
			go mergeCountGo(A, p, q, a)
			go mergeCountGo(A, q, r, b)
			c += <-a
			c += <-b
		} else {
			c += MergeCountGo(A, p, q)
			c += MergeCountGo(A, q, r)
		}
		c += count(A, p, q, r)
	}
	return c
}
