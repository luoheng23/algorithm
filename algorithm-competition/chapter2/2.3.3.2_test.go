package chapter2

import "testing"

func TestCombine(t *testing.T) {
	n, m, M, a := 3, 3, 10000, []int{1, 2, 3}
	if combine(n, m, M, a) != 6 {
		t.Errorf("TestCombine failed. Expected 6, Got %d.", combine(n, m, M, a))
	}

}
