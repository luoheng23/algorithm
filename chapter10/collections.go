// include Stack, Queue, DoubleLinkedList
package chapter10

import "fmt"

type Stack struct {
	A   []int
	top int
	n   int
}

func CreateStack(n int) Stack {
	s := Stack{}
	s.A, s.top, s.n = make([]int, n), -1, n
	return s
}

func (s *Stack) IsEmpty() bool {
	return s.top == -1
}

func (s *Stack) Push(x int) error {
	if s.top+1 < s.n {
		s.top++
		s.A[s.top] = x
		return nil
	} else {
		return fmt.Errorf("overflow")
	}
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, fmt.Errorf("underflow")
	} else {
		s.top--
		return s.A[s.top+1], nil
	}
}

type Queue struct {
	A      []int
	tail   int
	head   int
	length int
}

func CreateQueue(n int) Queue {
	q := Queue{}
	q.A, q.tail, q.head, q.length = make([]int, n), 0, 0, n
	return q
}

func (q *Queue) EnQueue(x int) error {
	if q.head == (q.tail+1)%q.length {
		return fmt.Errorf("overflow")
	}
	q.A[q.tail] = x
	q.tail = (q.tail + 1) % q.length
	return nil
}

func (q *Queue) DeQueue() (int, error) {
	if q.head == q.tail {
		return 0, fmt.Errorf("underflow")
	}
	x := q.A[q.head]
	q.head = (q.head + 1) % q.length
	return x, nil
}

type node struct {
	key  int
	prev *node
	next *node
}

type DoubleLinkedList struct {
	NIL *node // sentinel
}

func CreateDoubleLinkedList() DoubleLinkedList {
	d := DoubleLinkedList{NIL: &node{}}
	d.NIL.prev, d.NIL.next = d.NIL, d.NIL
	return d
}

func (d *DoubleLinkedList) Search(k int) *node {
	x := d.NIL.next
	for x != d.NIL && x.key != k {
		x = x.next
	}
	return x
}

func (d *DoubleLinkedList) Delete(x *node) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

func (d *DoubleLinkedList) Insert(x *node) {
	x.next = d.NIL.next
	d.NIL.next.prev = x
	d.NIL.next = x
	x.prev = d.NIL
}
