package chapter2

import "testing"

func TestMergeCount(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	mergeS := MergeCount(s[:], 0, len(s))
	countS := 5 + 3 + 5 + 5 + 3 + 3 + 2 + 1
	if mergeS != countS {
		t.Errorf("MergeCount failed. Expected %v, Got %v", countS, mergeS)
	}
}

func TestMergeCountGo(t *testing.T) {
	s := [10]int{5, 3, 7, 8, 4, 6, 0, 9, 2, 1}
	mergeS := MergeCountGo(s[:], 0, len(s))
	countS := 5 + 3 + 5 + 5 + 3 + 3 + 2 + 1
	if mergeS != countS {
		t.Errorf("MergeCount failed. Expected %v, Got %v", countS, mergeS)
	}
}
