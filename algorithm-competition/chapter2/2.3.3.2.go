package chapter2

// 多重组合数
func combine(n, m, M int, a []int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j-1-a[i] >= 0 {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j] - dp[i][j-1-a[i]] + M) % M
			} else {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j]) % M
			}
		}
	}
	return dp[n][m]
}
