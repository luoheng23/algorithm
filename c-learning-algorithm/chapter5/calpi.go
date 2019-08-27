package chapter5

import (
	"math"
	"math/rand"
	"time"
)

// CalPi calculate pi
// it takes a long time, don't use it
func CalPi() float64 {
	rand.Seed(time.Now().UnixNano())
	iter := int(math.Pow(10, 8))
	n := 0
	for i := 0; i < iter; i++ {
		if math.Pow(rand.Float64(), 2)+math.Pow(rand.Float64(), 2) < 1 {
			n++
		}
	}
	return 4 * float64(n) / float64(iter)
}
