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

// BinarySearch for value in sorted A
func BinarySearch(A []int, value int) int {
	low, high := 0, len(A)-1
	for low < high {
		mid := (low + high) / 2
		if A[mid] > value {
			high = mid
		} else if A[mid] < value {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
