package chapter3

import "testing"

func TestFaceTheRightWay(t *testing.T) {
	L := []int{1, 1, 0, 1, 0, 1, 1}
	K, M := faceTheRightWay(L)
	if K != 3 || M != 3 {
		t.Errorf("TestFaceTheRightWay failed. Expected (3,3), Got (%d, %d).", K, M)
	}

}
