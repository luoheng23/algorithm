package main

import "fmt"

func isPopSeries(push []int, pop []int, n int) bool {
	stack := make([]int, 0, len(push))
	i1, i2 := 0, 0
	for i2 < n {
		if len(stack) == 0 || stack[len(stack)-1] != pop[i2] {
			if i1 < n {
				stack = append(stack, push[i1])
				i1++
			} else {
				return false
			}
			for len(stack) != 0 && stack[len(stack)-1] == pop[i2] {
				stack = stack[:len(stack)-1]
				i2++
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isPopSeries([]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, 5))
	fmt.Println(isPopSeries([]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, 5))
}