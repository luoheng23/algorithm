package chapter2

// n以内的素数数量
func primeN(n int) int {
	var prime int
	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			prime++
		}
		for j := 2 * i; j <= n; j += i {
			isPrime[j] = false
		}
	}
	return prime
}
