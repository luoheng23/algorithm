package chapter2

import "testing"

func TestTotalBag(t *testing.T) {
	weight, value, totalW := []int{3, 4, 2}, []int{4, 5, 3}, 7
	if totalBag(weight, value, totalW) != 10 {
		t.Errorf("TestBag01 failed. Expected 7, Got %d.", totalBag(weight, value, totalW))
	}

}
