package chapter1

import "sort"

func binarySearch(s int, sum2 []int) bool {
	l, r := 0, len(sum2)
	for r-l >= 1 {
		mid := (r + l) / 2
		if sum2[mid] == s {
			return true
		} else if sum2[mid] < s {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}

// 四个数字的和为m，可重复
func sumM(n, m int, x []int) bool {
	sum2 := make([]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum2[i*n+j] = x[i] + x[j]
		}
	}
	sort.Ints(sum2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if binarySearch(m-x[i]-x[j], sum2) {
				return true
			}
		}
	}
	return false

}
