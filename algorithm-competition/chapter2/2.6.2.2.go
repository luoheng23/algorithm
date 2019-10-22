package chapter2

import "math"

func primeAB(a, b int64) int {
	prime := make([]bool, b-a)
	isPrimeSmall := make([]bool, int64(math.Sqrt(float64(b)))+1)
	for i := int64(2); i*i < b; i++ {
		if !isPrimeSmall[i] {
			for j := 2 * i; j*j < b; j += i {
				isPrimeSmall[j] = true
			}
			for j := int64(Max(2, int((a+i-1)/i))) * i; j < b; j += i {
				prime[j-a] = true
			}
		}
	}
	sum := 0
	for i := 0; i < len(prime); i++ {
		if !prime[i] {
			sum++
		}
	}
	return sum
}
