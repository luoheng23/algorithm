package chapter2

// 01背包问题
func bag01(weight []int, value []int, totalW int) int {
	dp := make([]int, totalW+1)
	for i := 0; i < len(weight); i++ {
		for j := totalW; j >= 0; j-- {
			if j >= weight[i] && dp[j-weight[i]]+value[i] > dp[j] {
				dp[j] = dp[j-weight[i]] + value[i]
			}
		}
	}
	return dp[totalW]
}
