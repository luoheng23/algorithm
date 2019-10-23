package chapter3

import "sort"

func c(x float64, weight, value []float64, k int) bool {
	ans := make([]float64, len(weight))
	for i := 0; i < len(weight); i++ {
		ans[i] = value[i] - x*weight[i]
	}
	sort.Float64s(ans)
	var sum float64
	for i := 0; i < k; i++ {
		sum += ans[len(ans)-i-1]
	}
	return sum >= 0
}

// 最大化平均值
func maxAve(weight, value []float64, k int) float64 {
	s, e := 0.0, float64(^uint(0)>>1)
	for s+.001 < e {
		m := (s + e) / 2
		if c(m, weight, value, k) {
			s = m
		} else {
			e = m
		}
	}
	return s
}
