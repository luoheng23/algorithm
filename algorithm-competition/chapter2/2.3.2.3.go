package chapter2

func lis(a []int) int {
	dp := make([]int, len(a))
	for i := 0; i < len(dp); i++ {
		dp[i] = int(^uint(0) >> 1)
	}
	for i := 0; i < len(a); i++ {
		s, e := 0, len(dp)
		for s < e {
			m := (s + e) / 2
			if a[i] > dp[m] {
				s = m + 1
			} else if a[i] < dp[m] {
				e = m
			} else {
				break
			}
		}
		if s == e {
			dp[s] = a[i]
		}
	}
	for i := 0; i < len(dp); i++ {
		if dp[i] == int(^uint(0)>>1) {
			return i
		}
	}
	return 0
}
