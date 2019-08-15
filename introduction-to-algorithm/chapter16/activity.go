package chapter16

// GreedyActivitySelector O(n)
// f is sorted
func GreedyActivitySelector(s, f []int) []int{
	n := len(s)
	A := []int{0}
	k := 0
	for m := 1; m < n; m++ {
		if s[m] >= f[k] {
			A = append(A, m)
			k = m
		}
	}
	return A
}