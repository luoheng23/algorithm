package chapter3

import "testing"

func TestCow(t *testing.T) {
	a, M := []int{1, 2, 8, 4, 9}, 3
	if aggressiveCows(a, M) != 3 {
		t.Errorf("TestCow failed. Expected 3, Got %d.", aggressiveCows(a, M))
	}
}
