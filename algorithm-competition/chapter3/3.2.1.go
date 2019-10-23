package chapter3

func subMinLen(seq []int, S int) int {
	res := len(seq) + 1
	s, t, sum := 0, 0, 0
	for {
		for t < len(seq) && sum < S {
			sum += seq[t]
			t++
		}
		if sum < S {
			break
		}
		for s < t && sum > S {
			sum -= seq[s]
			s++
		}
		if res > t-s+1 {
			res = t - s + 1
		}
	}
	if res > len(seq) {
		return 0
	}
	return res
}
