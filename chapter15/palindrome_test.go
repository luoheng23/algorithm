package chapter15

import (
	"strings"
	"testing"
)

func TestPalindromeSlow(t *testing.T) {
	s1 := convert("caracter")
	s2 := convert("ttccabbaccd")
	if unConvert(PalindromeSlow(s1)) != "carac" || unConvert(PalindromeSlow(s2)) != "ccabbacc" {
		t.Errorf("PalindromeSlow failed. Expected (carac, ccabbacc), Got (%s, %s)", unConvert(PalindromeSlow(s1)), unConvert(PalindromeSlow(s2)))
	}
}

func BenchmarkPalindromeSlow(b *testing.B) {
	s := "abcdefafadsafasdfsdfsccdfhzhvxzjkhfkfhskdhfjsfewufiwe"
	s = convert(strings.Repeat(s, 200) + "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	for i := 0; i < b.N; i++ {
		PalindromeSlow(s)
	}
}

func TestPalindromeFast(t *testing.T) {
	s1 := convert("caracter")
	s2 := convert("ttccabbaccd")
	if unConvert(PalindromeFast(s1)) != "carac" || unConvert(PalindromeFast(s2)) != "ccabbacc" {
		t.Errorf("PalindromeFast failed. Expected (carac, ccabbacc), Got (%s, %s)", unConvert(PalindromeFast(s1)), unConvert(PalindromeFast(s2)))
	}
}

func BenchmarkPalindromeFast(b *testing.B) {
	s := "abcdefafadsafasdfsdfsccdfhzhvxzjkhfkfhskdhfjsfewufiwe"
	s = convert(strings.Repeat(s, 200) + "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa") 
	for i := 0; i < b.N; i++ {
		PalindromeFast(s)
	}
}

func TestPalindromeVeryFast(t *testing.T) {
	// s1 := "caracter"
	// s2 := "ttccabbaccd"
	// s1 := "abcdefghij"
	s2 := convert("abcbaba")
	s1 := convert("abacabacaba")
	if unConvert(PalindromeVeryFast(s1)) != "abacabacaba" || unConvert(PalindromeVeryFast(s2)) != "abcba" {
		t.Errorf("PalindromeVeryFast failed. Expected (abacabacaba, aa), Got (%s, %s)", unConvert(PalindromeVeryFast(s1)), unConvert(PalindromeVeryFast(s2)))
	}
}

func BenchmarkPalindromeVeryFast(b *testing.B) {
	s := "abcdefafadsafasdfsdfsccdfhzhvxzjkhfkfhskdhfjsfewufiwe"
	s = convert(strings.Repeat(s, 200) + "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	for i := 0; i < b.N; i++ {
		PalindromeVeryFast(s)
	}
}
