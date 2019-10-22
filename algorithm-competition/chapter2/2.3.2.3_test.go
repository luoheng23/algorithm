package chapter2

import "testing"

func TestLis(t *testing.T) {
	a := []int{4, 2, 3, 1, 5}
	if lis(a) != 3 {
		t.Errorf("TestList failed. Expected 3, Got %d.", lis(a))
	}
}
