package chapter3

import (
	"math"
	"testing"
)

func TestCutRope(t *testing.T) {
	L, K := []float64{8.02, 7.43, 4.57, 5.39}, 11
	if math.Abs(cutRope(L, K)-2.00) > 0.01 {
		t.Errorf("TestCutRope failed. Expected 2.00, Got %f.", cutRope(L, K))
	}

}
