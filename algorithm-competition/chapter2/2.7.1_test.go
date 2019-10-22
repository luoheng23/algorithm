package chapter2

import "testing"

func TestMinProduct(t *testing.T) {
	v1, v2 := []int{1, 3, -5}, []int{-2, 4, 1}
	if minProduct(v1, v2) != -25 {
		t.Errorf("TestMinProduct failed. Expected -25, Got %d.", minProduct(v1, v2))
	}
}
