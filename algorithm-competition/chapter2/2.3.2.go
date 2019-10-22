package chapter2

// 完全背包问题
func totalBag(weight, value []int, totalW int) int {
	dp := make([]int, totalW+1)
	for i := 0; i < len(weight); i++ {
		for j := weight[i]; j <= totalW; j++ {
			dp[j] = Max(dp[j], dp[j-weight[i]]+value[i])
		}
	}
	return dp[totalW]
}
