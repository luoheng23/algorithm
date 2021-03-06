package chapter4

func maxCrossSubArray(A []int, low, mid, high int) (leftPos int, rightPos int, sum int) {
	left := -int(^uint32(0) >> 1)
	right := left
	leftPos, rightPos = -1, -1
	for i := mid; i > low; i-- {
		sum += A[i]
		if sum > left {
			left, leftPos = sum, i
		}
	}
	sum = 0
	for i := mid + 1; i < high; i++ {
		sum += A[i]
		if sum > right {
			right, rightPos = sum, i
		}
	}
	sum = left + right
	return leftPos, rightPos, sum
}

// MaxSubArray O(nlgn)
func MaxSubArray(A []int, low, high int) (leftPos int, rightPos int, sum int) {
	if high == low+1 {
		return low, high, A[low]
	}
	mid := (low + high) / 2
	leftLeftPos, leftRightPos, leftSum := MaxSubArray(A, low, mid)
	rightLeftPos, rightRightPos, rightSum := MaxSubArray(A, mid, high)
	crossLeftPos, crossRightPos, crossSum := maxCrossSubArray(A, low, mid, high)
	if leftSum >= rightSum && leftSum >= crossSum {
		return leftLeftPos, leftRightPos, leftSum
	} else if rightSum >= leftSum && rightSum >= crossSum {
		return rightLeftPos, rightRightPos, rightSum
	}
	return crossLeftPos, crossRightPos, crossSum
}

// MaxSubArrayLine O(n)
func MaxSubArrayLine(A []int) (leftPos int, rightPos int, sum int) {
	var l, r, s int
	for i, e := range A {
		if s >= 0 {
			s += e
			r++
		} else {
			s = e
			l, r = i, i
		}
		if sum < s {
			leftPos, rightPos, sum = l, r, s
		}
	}
	return leftPos, rightPos, sum
}
