package chapter5

import "testing"

func TestMyRand(t *testing.T) {
	a, b := 0, 10
	for i := 0; i < 10; i++ {
		rand := MyRand(a, b)
		if rand < a && rand > b {
			t.Errorf("MyRand failed. Expected (%d, %d), Got %d", a, b, rand)
		}
	}
}
