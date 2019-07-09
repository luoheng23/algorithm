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
	s = strings.Repeat(s, 200)
	for i := 0; i < b.N; i++ {
		PalindromeSlow(s)
	}
}

func TestPalindromeFast(t *testing.T) {
	s1 := "caracter"
	s2 := "ttccabbaccd"
	if PalindromeFast(s1) != "carac" || PalindromeFast(s2) != "ccabbacc" {
		t.Errorf("PalindromeFast failed. Expected (carac, ccabbacc), Got (%s, %s)", PalindromeFast(s1), PalindromeFast(s2))
	}
}

func BenchmarkPalindromeFast(b *testing.B) {
	s := "abcdefafadsafasdfsdfsccdfhzhvxzjkhfkfhskdhfjsfewufiwe"
	s = strings.Repeat(s, 200)
	for i := 0; i < b.N; i++ {
		PalindromeFast(s)
	}
}

// func TestPalindromeVeryFast(t *testing.T) {
// 	// s1 := "caracter"
// 	// s2 := "ttccabbaccd"
// 	// s1 := "abcdefghij"
// 	s2 := "aabcdefgc"
// 	s1 := "abacabacaba"
// 	if PalindromeVeryFast(s1) != "abacabacaba" || PalindromeVeryFast(s2) != "aa" {
// 		t.Errorf("PalindromeVeryFast failed. Expected (abacabacaba, aa), Got (%s, %s)", PalindromeVeryFast(s1), PalindromeVeryFast(s2))
// 	}
// }

// func BenchmarkPalindromeVeryFast(b *testing.B) {
// 	s := "abcdefafadsafasdfsdfsccdfhzhvxzjkhfkfhskdhfjsfewufiwe"
// 	s = strings.Repeat(s, 200)
// 	for i := 0; i < b.N; i++ {
// 		PalindromeVeryFast(s)
// 	}
// }
