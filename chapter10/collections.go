// Package chapter10 include Stack, Queue, DoubleLinkedList, tree
package chapter10

import "fmt"

// Stack stack
type Stack struct {
	A   []int
	top int
	n   int
}

// CreateStack init a stack
func CreateStack(n int) Stack {
	s := Stack{}
	s.A, s.top, s.n = make([]int, n), -1, n
	return s
}

func (s *Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) isFull() bool {
	return s.top + 1 == s.n
}

// Push an element
func (s *Stack) Push(x int) error {
	if !s.isFull() {
		s.top++
		s.A[s.top] = x
		return nil
	}
	return fmt.Errorf("overflow")
}

// Pop an element
func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, fmt.Errorf("underflow")
	} 
	s.top--
	return s.A[s.top+1], nil
}

// Queue queue
type Queue struct {
	A      []int
	tail   int
	head   int
	length int
}

// CreateQueue init a queue
func CreateQueue(n int) Queue {
	q := Queue{}
	q.A, q.tail, q.head, q.length = make([]int, n), 0, 0, n
	return q
}

func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) isFull() bool {
	return q.head == (q.tail + 1) % q.length
}

// EnQueue an element
func (q *Queue) EnQueue(x int) error {
	if q.isFull() {
		return fmt.Errorf("overflow")
	}
	q.A[q.tail] = x
	q.tail = (q.tail + 1) % q.length
	return nil
}

// DeQueue an element
func (q *Queue) DeQueue() (int, error) {
	if q.isEmpty() {
		return 0, fmt.Errorf("underflow")
	}
	x := q.A[q.head]
	q.head = (q.head + 1) % q.length
	return x, nil
}

// Node for linkedList
type Node struct {
	key  int
	prev *Node
	next *Node
}

// DoubleLinkedList linkedList
type DoubleLinkedList struct {
	NIL *Node // sentinel
}

// CreateDoubleLinkedList init a linkedList
func CreateDoubleLinkedList() DoubleLinkedList {
	d := DoubleLinkedList{NIL: &Node{}}
	d.NIL.prev, d.NIL.next = d.NIL, d.NIL
	return d
}

// Search for a Node
func (d *DoubleLinkedList) Search(k int) *Node {
	x := d.NIL.next
	for x != d.NIL && x.key != k {
		x = x.next
	}
	return x
}

// Delete a Node
func (d *DoubleLinkedList) Delete(x *Node) {
	x.prev.next = x.next
	x.next.prev = x.prev
}

// Insert a Node
func (d *DoubleLinkedList) Insert(x *Node) {
	x.next = d.NIL.next
	d.NIL.next.prev = x
	d.NIL.next = x
	x.prev = d.NIL
}

// TreeNode is the node of tree
type TreeNode struct {
	key int
	left, right *TreeNode
	p *TreeNode
}

// TreeNode2 left-child right-sibling representation
type TreeNode2 struct {
	key int
	leftChild, rightSibling *TreeNode2
	p *TreeNode
}

// Tree represents a tree
type Tree struct {
	root TreeNode
}

