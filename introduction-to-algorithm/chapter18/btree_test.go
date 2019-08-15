package chapter18

import (
	"fmt"
	"testing"
)

type Node struct {
	s *BNode
	i int
}

// Queue queue
type Queue struct {
	A      []Node
	tail   int
	head   int
	length int
}

// CreateQueue init a queue
func CreateQueue(n int) Queue {
	q := Queue{}
	q.A, q.tail, q.head, q.length = make([]Node, n), 0, 0, n
	return q
}

func (q *Queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) isFull() bool {
	return q.head == (q.tail+1)%q.length
}

// EnQueue an element
func (q *Queue) EnQueue(x Node) error {
	if q.isFull() {
		return fmt.Errorf("overflow")
	}
	q.A[q.tail] = x
	q.tail = (q.tail + 1) % q.length
	return nil
}

// DeQueue an element
func (q *Queue) DeQueue() (Node, error) {
	if q.isEmpty() {
		return Node{}, fmt.Errorf("underflow")
	}
	x := q.A[q.head]
	q.head = (q.head + 1) % q.length
	return x, nil
}

func (b *BTree) printInOrder() {
	queue := CreateQueue(50)
	i, cur := 0, 0
	queue.EnQueue(Node{b.root, 0})
	fmt.Println("start: ")
	for !queue.isEmpty() {
		n, _ := queue.DeQueue()
		if n.i == cur {
			fmt.Print("{", n.s, n.i, "}")
		} else {
			cur++
			fmt.Println("")
			fmt.Print("{", n.s, n.i, "}")
		}
		for _, c := range n.s.c {
			if c != nil {
				queue.EnQueue(Node{c, n.i + 1})
				i++
			}
		}
	}
	fmt.Println("\nend")
}

func TestBTree(t *testing.T) {
	b := CreateTree(2)
	s := []rune{'F', 'S', 'Q', 'K', 'C', 'L', 'H', 'T', 'V', 'W'}
	for _, r := range s {
		b.Insert(r)
	}
	b.printInOrder()
	b.Delete(b.root, 'F')
	b.Delete(b.root, 'H')
	b.Delete(b.root, 'K')
	b.printInOrder()
}
