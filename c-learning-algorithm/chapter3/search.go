package chapter3

// Search for value in A
func Search(A []int, value int) int {
	for i, a := range A {
		if value == a {
			return i
		}
	}
	return -1
}
