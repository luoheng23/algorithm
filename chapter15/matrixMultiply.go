package chapter15

import "fmt"

// MatrixChainOrder minimize times of multiply
func MatrixChainOrder(p []int) ([][]int, [][]int){
	const INTMAX = int(^uint32(0) >> 1)
	n := len(p) - 1
	m := make([][]int, n)
	s := make([][]int, n-1)
	for i := 0; i<n; i++ {
		m[i] = make([]int, n)
	}
	for i:= 0; i<n-1; i++ {
		s[i] = make([]int, n-1)
	}
	for l := 2; l < n + 1; l++ {
		for i := 1; i <= n - l + 1; i++ {
			j := i + l - 1
			m[i-1][j-1] = INTMAX
			for k := i; k <= j - 1; k++ {
				q := m[i-1][k-1] + m[k][j-1] + p[i-1]*p[k]*p[j]
				if q < m[i-1][j-1] {
					m[i-1][j-1] = q
					s[i-1][j-2] = k
				}
			}
		}
	}
	return m, s
}

// PrintOptimalParens result
func PrintOptimalParens(s [][]int, i, j int) {
	if i == j {
		fmt.Print("A_", i)
	} else {
		fmt.Print("(")
		PrintOptimalParens(s, i, s[i-1][j-2])
		PrintOptimalParens(s, s[i-1][j-2]+1, j)
		fmt.Print(")")
	}
}