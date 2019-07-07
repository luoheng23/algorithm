package chapter4

import "testing"

func TestMaxSubArray(t *testing.T) {
	var array = [16]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 28, -8, 7}
	left, right, sum := 7, 13, 44
	l, r, s := MaxSubArray(array[:], 0, len(array))
	if l != left || r != right || s != sum {
		t.Errorf("MaxSubArray failed. Expected (%d, %d, %d), Got (%d, %d, %d).", left, right, sum, l, r, s)
	}
}

func BenchmarkMaxSubArray(b *testing.B) {
	var array = [16]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 28, -8, 7}
	for i := 0; i < b.N; i++ {
		MaxSubArray(array[:], 0, len(array))
	}
}

func TestMaxSubArrayLine(t *testing.T) {
	var array = [16]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 28, -8, 7}
	left, right, sum := 7, 13, 44
	l, r, s := MaxSubArrayLine(array[:])
	if l != left || r != right || s != sum {
		t.Errorf("MaxSubArrayLine failed. Expected (%d, %d, %d), Got (%d, %d, %d).", left, right, sum, l, r, s)
	}
}
func BenchmarkMaxSubArrayLine(b *testing.B) {
	var array = [16]int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 28, -8, 7}
	for i := 0; i < b.N; i++ {
		MaxSubArrayLine(array[:])
	}
}
