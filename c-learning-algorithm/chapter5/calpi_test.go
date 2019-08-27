package chapter5

import (
	"math"
	"testing"
)

func TestCalPi(t *testing.T) {
	pi := CalPi()
	if math.Abs(pi-math.Pi) > 0.001 {
		t.Errorf("CalPi failed. Expected %f, Got %f.", math.Pi, pi)
	}
}
