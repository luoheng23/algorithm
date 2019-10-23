package chapter3

import "testing"

func TestSubMinLen(t *testing.T) {
	a, S := []int{1, 2, 3, 4, 5}, 11
	if subMinLen(a, S) != 3 {
		t.Errorf("TestSubMinLen failed. Expected 3, Got %d.", subMinLen(a, S))
	}

}
