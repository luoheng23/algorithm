package chapter2

// 硬币问题
func coin(A int) int {
	coins := [6]int{1, 5, 10, 50, 100, 500}
	ans := 0
	for i := len(coins) - 1; i >= 0; i-- {
		t := A / coins[i]
		A -= t * coins[i]
		ans += t
	}
	return ans
}
