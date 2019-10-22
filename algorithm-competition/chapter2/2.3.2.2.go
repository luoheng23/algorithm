package chapter2

// 多重部分和
func sumK(a []int, m []int, K int) bool {
	dp := make([]int, K+1)
	for i := 0; i < len(a); i++ {
		for j := 0; j <= K; j++ {
			if dp[j] >= 0 {
				dp[j] = m[i]
			} else if j < a[i] || dp[j-a[i]] <= 0 {
				dp[j] = -1
			} else {
				dp[j] = dp[j-a[i]] - 1
			}
		}
	}
	return dp[K] >= 0
}
