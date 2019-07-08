package chapter9

func Min(A []int) int {
	if len(A) < 1 {
		return 0
	}
	min := A[0]
	for _, a := range A {
		if a < min {
			min = a
		}
	}
	return min
}

func Max(A []int) int {
	if len(A) < 1 {
		return 0
	}
	max := A[0]
	for _, a := range A {
		if a > max {
			max = a
		}
	}
	return max
}
