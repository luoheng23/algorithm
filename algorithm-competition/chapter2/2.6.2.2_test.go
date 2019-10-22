package chapter2

import "testing"

func TestPrimeAB(t *testing.T) {
	if primeAB(22, 37) != 3 || primeAB(22801763489, 22801787297) != 1000 {
		t.Errorf("TestPrimeAB failed. Expected (3,1000), Got (%d,%d).", primeAB(22, 37), primeAB(22801763489, 22801787297))
	}
}
