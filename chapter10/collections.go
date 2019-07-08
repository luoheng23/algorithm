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

func (s *stack) Push(x int) error {
	if s.top+1 < s.n {
		s.top++
		s.A[s.top] = x
		return nil
	} else {
		return fmt.Errorf("overflow")
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

type queue struct {
	A      []int
	tail   int
	head   int
	length int
}

func CreateQueue(n int) queue {
	q := queue{}
	q.A, q.tail, q.head, q.length = make([]int, n), 0, 0, n
	return q
}

func (q *queue) EnQueue(x int) error {
	if q.head == (q.tail+1)%q.length {
		return fmt.Errorf("overflow")
	}
	q.A[q.tail] = x
	q.tail = (q.tail + 1) % q.length
	return nil
}

func (q *queue) DeQueue() (int, error) {
	if q.head == q.tail {
		return 0, fmt.Errorf("underflow")
	}
	x := q.A[q.head]
	q.head = (q.head + 1) % q.length
	return x, nil
}
