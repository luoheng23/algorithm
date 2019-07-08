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

func MinMax(A []int) (min int, max int) {
	if len(A) < 1 {
		return 0, 0
	}
	min, max = A[0], A[0]
	mi, ma := 0, 0
	for i := 1; i < len(A)-1; i += 2 {
		if A[i] > A[i+1] {
			mi, ma = A[i+1], A[i]
		} else {
			mi, ma = A[i], A[i+1]
		}
		if ma > max {
			max = ma
		}
		if mi < min {
			min = mi
		}
	}
	return min, max
}
