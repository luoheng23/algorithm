package chapter2

import "sort"

// 向量的最小乘积
func minProduct(v1, v2 []int) int64 {
	sort.Ints(v1)
	sort.Ints(v2)
	var ans int64
	for i := 0; i < len(v1); i++ {
		ans += int64(v1[i] * v2[len(v2)-1-i])
	}
	return ans
}
