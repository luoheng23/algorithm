package chapter2

import "testing"

func TestCoin(t *testing.T) {
	A := 620
	if coin(A) != 4 {
		t.Errorf("TestCoin failed. Expected 4, Got %d", coin(A))
	}
}
