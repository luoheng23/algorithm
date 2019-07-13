package chapter15

import "testing"

func TestBag01(t *testing.T) {
	bagWeight := [5]int{4, 4, 7, 6, 5}
	maxWeight := 8
	s := Bag01(bagWeight[:], maxWeight)
	if s != 8 {
		t.Errorf("Test Bag01 failed. Expected 8, Got %d", s)
	}
}
