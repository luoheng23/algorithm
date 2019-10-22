package chapter3

// 最小i，二分
func minI(a []int, k int) int {
	s, e := 0, len(a)
	for s < e {
		m := (s + e) / 2
		if a[m] >= k {
			e = m
		} else {
			s = m + 1
		}
	}
	return s
}
