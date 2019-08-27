package chapter3

// ShellSort performs the shell sort
func ShellSort(A []int, incr int) {
	for incr > 0 {
		for i := incr; i < len(A); i++ {
			j := i - incr
			for j >= 0 {
				if A[j] > A[j+incr] {
					A[j], A[j+incr] = A[j+incr], A[j]
					j -= incr
				} else {
					break
				}
			}
		}
		incr /= 2
	}
}
