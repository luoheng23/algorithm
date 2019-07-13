package chapter5

import "testing"
import "math"

func TestMyRand(t *testing.T) {
	a, b := 0, 10
	for i := 0; i < 10; i++ {
		rand := MyRand(a, b)
		if rand < a && rand > b {
			t.Errorf("MyRand failed. Expected (%d, %d), Got %d", a, b, rand)
		}
	}
}

func TestMyBiasedRand(t *testing.T) {
	s := 100000
	t0, t1 := 0, 0
	for i := 0; i < s; i++ {
		if MyBiasedRand() == 0 {
			t0++
		} else {
			t1++
		}
	}
	if math.Abs(float64(t0-t1))/float64(s) >= 0.1 {
		t.Errorf("MyBiasedRand failed. Expected (50000, 50000), Got (%d, %d).", t0, t1)

	}
}
