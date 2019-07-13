package chapter8

// all elements are between 0 and k(include)
func CountSort(A []int, k int) {
	c := make([]int, k+1)
	for _, a := range A {
		c[a]++
	}
	for i := 1; i < k+1; i++ {
		c[i] += c[i-1]
	}
	b := make([]int, len(A))
	for j := len(A) - 1; j >= 0; j-- {
		c[A[j]]--
		b[c[A[j]]] = A[j]
	}
	for i, _ := range A {
		A[i] = b[i]
	}
}
