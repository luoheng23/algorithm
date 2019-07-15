package chapter15

// CutRodSlow solve cut rod problem
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

// CutRodMemorizedAux solve cut rod problem
func CutRodMemorizedAux(p []int, n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
	}
	q := 0
	for i := 1; i <= n; i++ {
		newQ := p[i] + CutRodMemorizedAux(p, n-i, r)
		if q < newQ {
			q = newQ
		}
	}
	r[n] = q
	return q
}

// CutRodMemorizedFast solve cut rod problem
func CutRodMemorizedFast(p []int, n int) int {
	r := make([]int, n+1)
	for i := 0; i <= n; i++ {
		r[i] = -1
	}
	return CutRodMemorizedAux(p, n, r)
}

// CutRodBottomFast solve cut rod problem
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

func cutRodSolution(p []int, n int) (r, s []int) {
	r, s = make([]int, n+1), make([]int, n+1)
	for i := 1; i <= n; i++ {
		q := -1
		for j := 1; j <= i; j++ {
			newQ := p[j] + r[i-j]
			if q < newQ {
				q = newQ
				s[i] = j
			}
		}
		r[i] = q
	}
	return r, s
}

// CutRodSolution solve cut rod problem
func CutRodSolution(p []int, n int) []int {
	_, s := cutRodSolution(p, n)
	r := []int{}
	for n > 0 {
		r = append(r, s[n])
		n -= s[n]
	}
	return r
}
