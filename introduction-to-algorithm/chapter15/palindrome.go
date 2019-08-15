package chapter15

import "strings"

// add # between strings
func convert(p string) string {
	return "#" + strings.Join(strings.Split(p, ""), "#") + "#"
}

func unConvert(p string) string {
	return strings.Join(strings.Split(p, "#"), "")
}

// PalindromeSlow O(n^2)
func PalindromeSlow(p string) string {
	length := len(p)
	maxI, maxL := 0, 0
	for i := 0; i < length; i++ {
		j, k := i-1, i+1
		for j >= 0 && k < length && p[j] == p[k] {
			j--
			k++
		}
		if maxL < k-j-1 {
			maxI = i
			maxL = k - j - 1
		}
	}
	return p[maxI-maxL/2 : maxI+maxL/2+1]
}

// PalindromeFast O(n)
func PalindromeFast(p string) string {
	length := len(p)
	s := make([]int, length)
	id, mx := 0, 0
	for i := 0; i < length; i++ {
		if mx > i && mx-i > s[2*id-i] {
			s[i] = s[2*id-i]
			continue
		}
		if mx-i > 1 {
			s[i] = mx - i
		} else {
			s[i] = 1
		}
		for i-s[i] >= 0 && i+s[i] < length && p[i+s[i]] == p[i-s[i]] {
			s[i]++
		}
		if s[i]+i > mx {
			mx, id = s[i]+i, i
		}
	}
	maxI, maxL := 0, 0
	for i, a := range s {
		if a > maxL {
			maxI, maxL = i, a
		}
	}
	return p[maxI-maxL+1 : maxI+maxL]
}

// PalindromeVeryFast dynamic programming O(n)
func PalindromeVeryFast(p string) string {
	length := len(p)
	maxI, maxL, id := 0, 0, 0
	for i := 1; i < length; i++ {
		for 2*id-i < 0 || p[2*id-i] != p[i] {
			id++
		}
		if i-id > maxL {
			maxI, maxL = id, i-id
		}
	}
	return p[maxI-maxL : maxI+maxL+1]
}
