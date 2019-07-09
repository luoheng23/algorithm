package chapter15

import (
	"strings"
	"testing"
)

func TestPalindromeSlow(t *testing.T) {
	s1 := "caracter"
	s2 := "ttccabbaccd"
	if PalindromeSlow(s1) != "carac" || PalindromeSlow(s2) != "ccabbacc" {
		t.Errorf("PalindromeSlow failed. Expected (carac, ccabbacc), Got (%s, %s)", PalindromeSlow(s1), PalindromeSlow(s2))
	}
}

func BenchmarkPalindromeSlow(b *testing.B) {
	s := "abcdefafadsafasdfsdfsccdfhzhvxzjkhfkfhskdhfjsfewufiwe"
	s = strings.Repeat(s, 20)
	for i := 0; i < b.N; i++ {
		PalindromeSlow(s)
	}
}
