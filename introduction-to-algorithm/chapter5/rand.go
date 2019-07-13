package chapter5

import (
	"math/rand"
	"time"
)

func MyRand(a, b int) int {
	r := rand.New(rand.NewSource(99))
	return int(r.Float64()*float64(b-a)) + a
}

// here rand(0, 1) is a biased rand, p to 1, 1-p to 0
// I will make it 0.5 to 1, 0.5 to 0
// p(1-p) means 1, (1-p)p means 0
func MyBiasedRand() int {
	for {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		a, b := r.Intn(2), r.Intn(2)
		if a == 1 && b == 0 {
			return 1
		} else if a == 0 && b == 1 {
			return 0
		}
	}
}
