package chapter15

import (
	"testing"
)

func TestMatrixChainOrder(t *testing.T) {
	p := []int{30, 35, 15, 5, 10, 20, 25}
	_, s := MatrixChainOrder(p)
	PrintOptimalParens(s, 1, len(s))
}
