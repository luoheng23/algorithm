package chapter2

import "testing"

func TestBestCowLine(t *testing.T) {
	s := "ACDBCB"
	if bestCowLine(s) != "ABCBCD" {
		t.Errorf("bestCowLine failed. Expected ABCBCD, Got %s", bestCowLine(s))
	}
}
