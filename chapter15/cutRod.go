package chapter15

func CutRodSlow(p []int, n int) int {
	if n == 0 {
		return 0
	}
	q := -1
	for i := 1; i <= n; i++ {
		newQ := p[i] + CutRodSlow(p, n-i)
		if q < newQ {
			q = newQ
		}
	}
	return q
}

func CutRodMemorizedAux(p []int, n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
	}
	q := 0
	if n != 0 {
		for i := 1; i <= n; i++ {
			newQ := p[i] + CutRodMemorizedAux(p, n-i, r)
			if q < newQ {
				q = newQ
			}
		}
	}
	r[n] = q
	return q
}

func CutRodMemorizedFast(p []int, n int) int {
	r := make([]int, n+1)
	for i := 0; i <= n; i++ {
		r[i] = -1
	}
	return CutRodMemorizedAux(p, n, r)
}

func CutRodBottomFast(p []int, n int) int {
	r := make([]int, n+1)
	for i := 1; i <= n; i++ {
		q := -1
		for j := 1; j <= i; j++ {
			newQ := p[j] + r[i-j]
			if q < newQ {
				q = newQ
			}
		}
		r[i] = q
	}
	return r[n]
}
