package chapter2

import "testing"

func TestPrimeN(t *testing.T) {
	if primeN(11) != 5 || primeN(1000000) != 78498 {
		t.Errorf("TestPrimeN failed. Expected (5, 78498), Got (%d, %d).", primeN(11), primeN(1000000))
	}
}
