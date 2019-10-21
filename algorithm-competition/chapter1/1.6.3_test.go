package chapter1

import "testing"

func TestSumM(t *testing.T) {
	n, m, x := 3, 10, []int{1, 3, 5}
	if !sumM(n, m, x) {
		t.Error("SumM failed. Expected true, Got false")
	}
}
