package chapter15

import (
	"strings"
)

// add # between strings
func convert(p string) string {
	return "#" + strings.Join(strings.Split(p, ""), "#") + "#"
}

func unConvert(p string) string {
	return strings.Join(strings.Split(p, "#"), "")
}

func PalindromeSlow(p string) string {
	cP := convert(p)
	length := len(cP)
	maxI, maxL := 0, 0
	for i := 0; i < length; i++ {
		j, k := i, i
		for j >= 0 && k < length && cP[j] == cP[k] {
			j--
			k++
		}
		if maxL < k-j+1 {
			maxI = i
			maxL = k - j - 1
		}
	}
	return unConvert(cP[maxI-maxL/2 : maxI+maxL/2+1])
}

func PalindromeFast(p string) string {
	cP := convert(p)
	length := len(cP)
	s := make([]int, length)
	id, mx := 0, 0
	for i := 0; i < length; i++ {
		if mx > i {
			if mx-i >= s[2*id-i] {
				s[i] = s[2*id-i]
			} else {
				s[i] = mx - i
			}
		} else {
			s[i] = 1
		}
		for i-s[i] >= 0 && i+s[i] < length && cP[i+s[i]] == cP[i-s[i]] {
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
	return unConvert(cP[maxI-maxL+1 : maxI+maxL])

}

// func PalindromeVeryFast(p string) string {
// 	cP := convert(p)
// 	length := len(cP)
// 	maxI, maxL, lastI := 0, 0, 0
// 	for i := 1; i < length; i++ {
// 		if 2*lastI-i < 0 || cP[2*lastI-i] != cP[i] {
// 			if i-2 >= 0 && cP[i-2] == cP[i] {
// 				lastI = i - 1
// 			} else {
// 				lastI = i
// 			}
// 		}
// 		if i-lastI > maxL {
// 			maxI, maxL = lastI, i-lastI
// 		}
// 	}
// 	return unConvert(cP[maxI-maxL : maxI+maxL+1])
// }
