package chapter3

import (
	"math"
	"testing"
)

func TestMaxAve(t *testing.T) {
	k, weight, value := 2, []float64{2, 5, 2}, []float64{2, 3, 1}
	if math.Abs(maxAve(weight, value, k)-0.75) > 0.01 {
		t.Errorf("TestMaxAve failed. Expected 0.75, Got %f.", maxAve(weight, value, k))
	}

}
