package chapter3

import "testing"

func TestJessicaRead(t *testing.T) {
	a := []int{1, 8, 8, 8, 1}
	if jessicaRead(a) != 2 {
		t.Errorf("TestJessicaRead failed. Expected 2, Got %d.", jessicaRead(a))
	}

}
