package chapter8

// CountSortForRadixSort O(n)
func CountSortForRadixSort(A []int, d int) {
	dNum := 1
	for i := 0; i < d; i++ {
		dNum *= 10
	}
	c := make([]int, 11)
	for _, a := range A {
		c[a/dNum%10]++
	}
	for i := 1; i < 11; i++ {
		c[i] += c[i-1]
	}
	b := make([]int, len(A))
	for j := len(A) - 1; j >= 0; j-- {
		c[A[j]/dNum%10]--
		b[c[A[j]/dNum%10]] = A[j]
	}
	for i := range A {
		A[i] = b[i]
	}
}


// RadixSort O(d(n + k))
func RadixSort(A []int, d int) {
	for i := 1; i < d + 1; i++ {
		CountSortForRadixSort(A, i)
	}
}