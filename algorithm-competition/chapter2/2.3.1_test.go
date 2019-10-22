package chapter2

import "testing"

func TestBag01(t *testing.T) {
	weight, value, totalW := []int{2, 1, 3, 2}, []int{3, 2, 4, 2}, 5
	if bag01(weight, value, totalW) != 7 {
		t.Errorf("TestBag01 failed. Expected 7, Got %d.", bag01(weight, value, totalW))
	}

}
