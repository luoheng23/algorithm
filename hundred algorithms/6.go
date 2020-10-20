package main

import "fmt"

func count(arr []int, i int) int {
	cnt := 0
	for _, v := range arr {
		if v == i {
			cnt++
		}
	}
	return cnt
}

func work(arr []int) []int {
	var res = make([]int, len(arr))
	for {
		isChanged := false
		for i := range res {
			cnt := count(res, arr[i])
			if cnt != res[i] {
				res[i] = cnt
				isChanged = true
			}
		}
		if !isChanged {
			break
		}
	}
	return res
}

func main() {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(work(test[:]))
}