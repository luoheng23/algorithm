package chapter4

import "testing"

func TestComplexMultiply(t *testing.T) {
	comp1, comp2 := 3+4i, 5+6i
	if comp1*comp2 != Multiply(comp1, comp2) {
		t.Errorf("Multiply failed. Expected %.0f, Got %.0f", comp1*comp2, Multiply(comp1, comp2))
	}
}
