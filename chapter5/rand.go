package chapter5

import "math/rand"

func MyRand(a, b int) int {
	r := rand.New(rand.NewSource(99))
	return int(r.Float64()*float64(b-a)) + a
}
