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
