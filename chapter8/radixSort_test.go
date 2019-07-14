package chapter8

import "testing"

func TestRadixSort(t *testing.T) {
	s := [10]int{12345, 54321, 53124, 76438, 22333, 43221, 99999, 223, 44444, 4}
	sortedS := [10]int{4, 223, 12345, 22333, 43221, 44444, 53124, 54321, 76438, 99999}
	RadixSort(s[:], 5)
	if s != sortedS {
		t.Errorf("RadixSort failed. Expected %v, Got %v", sortedS, s)
	}
}
