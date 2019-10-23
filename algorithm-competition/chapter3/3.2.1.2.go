package chapter3

// 覆盖知识点
func jessicaRead(seq []int) int {
	all := make(map[int]bool)
	for i := 0; i < len(seq); i++ {
		all[seq[i]] = true
	}
	res, s, t, sum := len(seq), 0, 0, 0
	count := make(map[int]int)
	for {
		for t < len(seq) && sum < len(all) {
			if count[seq[t]] == 0 {
				sum++
			}
			count[seq[t]]++
			t++
		}
		if sum < len(all) {
			break
		}
		if res > t-s {
			res = t - s
		}
		count[seq[s]]--
		if count[seq[s]] == 0 {
			sum--
		}
		s++
	}
	return res
}
