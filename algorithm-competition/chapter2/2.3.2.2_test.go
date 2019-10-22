package chapter2

import "testing"

func TestSumK(t *testing.T) {
	a, m, K := []int{3, 5, 8}, []int{3, 2, 2}, 17
	if !sumK(a, m, K) {
		t.Errorf("TestSumK failed. Expected true, Got false.")
	}
}
