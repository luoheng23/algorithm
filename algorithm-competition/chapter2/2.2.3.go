package chapter2

import "bytes"

// 字典序最小问题
func bestCowLine(s string) string {
	a, b := 0, len(s)
	t := bytes.NewBufferString("")
	for a < b {
		left := false
		for i := a; i < b; i++ {
			if s[i] < s[b-a+i-1] {
				left = true
				break
			} else if s[i] > s[b-a+i-1] {
				break
			}
		}
		if left {
			t.WriteByte(s[a])
			a++
		} else {
			t.WriteByte(s[b-1])
			b--
		}
	}
	return t.String()
}
