package chapter3

func check(L []float64, m float64, K int) bool {
	num := 0
	for i := 0; i < len(L); i++ {
		num += int(L[i] / m)
	}
	return num >= K
}

// 切长度相同的绳子
func cutRope(L []float64, K int) float64 {
	var s, e float64
	s, e = 0, float64(^uint(0)>>1)
	for s+.001 < e {
		m := (s + e) / 2
		if check(L, m, K) {
			s = m
		} else {
			e = m
		}
	}
	return s
}
