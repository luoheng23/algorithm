package chapter8

import "testing"

func TestBucketSort(t *testing.T) {
	s := [10]float64{0.79, 0.13, 0.16, 0.64, 0.39, 0.20, 0.89, 0.53, 0.71, 0.42}
	sortedS := [10]float64{0.13, 0.16, 0.20, 0.39, 0.42, 0.53, 0.64, 0.71, 0.79, 0.89}
	BucketSort(s[:])
	if s != sortedS {
		t.Errorf("RadixSort failed. Expected %v, Got %v", sortedS, s)
	}
}
