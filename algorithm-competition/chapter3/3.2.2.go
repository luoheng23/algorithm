package chapter3

func calc(L []int, K int) int {
	res, sum := 0, 0
	f := make([]int, len(L))
	for i := 0; i+K <= len(L); i++ {
		if (L[i]+sum)%2 != 0 {
			res++
			f[i] = 1
		}
		sum += f[i]
		if i-K+1 >= 0 {
			sum -= f[i-K+1]
		}
	}
	for i := len(L) - K + 1; i < len(L); i++ {
		if (L[i]+sum)%2 != 0 {
			return -1
		}
		if i-K+1 >= 0 {
			sum -= f[i-K+1]
		}
	}
	return res
}

// 反转
func faceTheRightWay(L []int) (int, int) {
	K, M := 1, len(L)
	for i := 1; i <= len(L); i++ {
		m := calc(L, i)
		if m >= 0 && M > m {
			M, K = m, i
		}
	}
	return K, M
}
