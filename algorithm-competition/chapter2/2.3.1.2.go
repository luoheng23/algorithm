package chapter2

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

//最长公共子序列
func lcs(s, t string) int {
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(t)+1)
	}
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			}
			dp[i+1][j+1] = Max(dp[i+1][j+1], dp[i][j+1], dp[i+1][j])
		}
	}
	return dp[len(s)][len(t)]
}
