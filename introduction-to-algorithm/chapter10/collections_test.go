package chapter10

import "testing"

func TestStack(t *testing.T) {
	s := CreateStack(10)
	if v, ok := s.Pop(); ok == nil {
		t.Errorf("Stack failed. Expected error, Got %d", v)
	}
	if !s.isEmpty() {
		t.Errorf("Stack failed. Expected True, Got False")
	}
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	a, _ := s.Pop()
	b, _ := s.Pop()
	c, _ := s.Pop()
	d, _ := s.Pop()
	if a != 5 || b != 4 || c != 3 || d != 2 {
		t.Errorf("Stack failed. Expected (5, 4, 3, 2), Got (%d, %d, %d, %d)", a, b, c, d)
	}
	if !s.isEmpty() {
		t.Errorf("Stack failed. Expected True, Got False")
	}
}

func TestQueue(t *testing.T) {
	q := CreateQueue(10)
	if v, ok := q.DeQueue(); ok == nil {
		t.Errorf("Queue failed. Expected Error, Got %d", v)
	}
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)
	a, _ := q.DeQueue()
	b, _ := q.DeQueue()
	c, _ := q.DeQueue()
	d, _ := q.DeQueue()
	if a != 2 || b != 3 || c != 4 || d != 5 {
		t.Errorf("Queue failed. Expected (2, 3, 4, 5), Got (%d, %d, %d, %d)", a, b, c, d)
	}
	if v, ok := q.DeQueue(); ok == nil {
		t.Errorf("Queue failed. Expected Error, Got %d", v)
	}
}

func TestDoubleLinkedList(t *testing.T) {
	d := CreateDoubleLinkedList()
	d.Insert(&Node{key: 4})
	d.Insert(&Node{key: 6})
	d.Insert(&Node{key: 5})
	d.Insert(&Node{key: 3})
	for _, a := range [4]int{3, 4, 5, 6} {
		if d.Search(a).key != a {
			t.Errorf("DoubleLinkedList failed. Expected %d, Got %d", a, d.Search(a).key)
		}
	}
	d.Delete(d.Search(4))
	for _, a := range [3]int{3, 5, 6} {
		if d.Search(a).key != a {
			t.Errorf("DoubleLinkedList failed. Expected %d, Got %d", a, d.Search(a).key)
		}
	}
	if d.Search(4).key == 4 {
		t.Errorf("DoubleLinkedList failed. Expected %d, Got %d", 4, d.Search(4).key)
	}

}
