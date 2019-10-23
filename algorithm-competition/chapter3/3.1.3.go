package chapter3

import "sort"

func cow(L []int, d int, M int) bool {
	last := 0
	for i := 1; i < M; i++ {
		crt := last + 1
		for crt < len(L) && L[crt]-L[last] < d {
			crt++
		}
		if crt == len(L) {
			return false
		}
		last = crt
	}
	return true
}

// 放牛
func aggressiveCows(L []int, M int) int {
	sort.Ints(L)
	var s, e int
	s, e = 0, int(^uint(0)>>1)
	for s < e-1 {
		m := (s + e) / 2
		if cow(L, m, M) {
			s = m
		} else {
			e = m
		}
	}
	return s
}
