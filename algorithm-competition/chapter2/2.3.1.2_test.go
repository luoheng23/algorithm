package chapter2

import "testing"

func TestLcs(t *testing.T) {
	s1, s2 := "abcd", "becd"
	if lcs(s1, s2) != 3 {
		t.Errorf("TestLcs failed. Expected 3, Got %d.", lcs(s1, s2))
	}

}
