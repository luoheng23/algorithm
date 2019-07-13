package chapter15

import "testing"

func TestBag01(t *testing.T) {
	bagWeight := [5]int{4, 4, 7, 6, 5}
	bagValue := [5]int{5, 5, 9, 8, 7}
	maxWeight := 8
	s := Bag01(bagWeight[:], bagValue[:], maxWeight)
	if len(s) != 2 || s[0] != 4 || s[1] != 4 {
		t.Errorf("Test Bag01 failed. Expected (4, 4), Got %v", s)
	}
}
