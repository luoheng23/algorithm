package chapter2

import "testing"

func TestDivNM(t *testing.T) {
	n, m, M := 4, 3, 10000
	if divNM(n, m, M) != 4 {
		t.Errorf("TestDivNM failed. Expected 4, Got %d.", divNM(n, m, M))
	}
}
