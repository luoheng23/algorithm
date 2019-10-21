package chapter1

// Min return min of x
func Min(x ...int) int {
	min := x[0]
	for _, a := range x {
		if a < min {
			min = a
		}
	}
	return min
}

// Max return max of x
func Max(x ...int) int {
	max := x[0]
	for _, a := range x {
		if a > max {
			max = a
		}
	}
	return max
}

// 计算蚂蚁落下杆子所需的最短时间和最长时间
func ants(n, L int, x []int) (int, int) {
	min, max := 0, 0
	for i := 0; i < n; i++ {
		min = Max(min, Min(x[i], L-x[i]))
	}
	for i := 0; i < n; i++ {
		max = Max(max, Max(x[i], L-x[i]))
	}
	return min, max
}
