package chapter10

import "fmt"

type stack struct {
	A   []int
	top int
	n   int
}

func CreateStack(n int) stack {
	s := stack{}
	s.A, s.top, s.n = make([]int, n), -1, n
	return s
}

func (s *stack) IsEmpty() bool {
	return s.top == -1
}

func (s *stack) Push(x int) bool {
	if s.top+1 < s.n {
		s.top++
		s.A[s.top] = x
		return true
	} else {
		return false
	}
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("underflow")
	} else {
		s.top--
		return s.A[s.top+1], nil
	}
}
