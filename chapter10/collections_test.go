package chapter10

import "testing"

func TestStack(t *testing.T) {
	s := CreateStack(10)
	if v, ok := s.Pop(); ok == nil {
		t.Errorf("Stack failed. Expected error, Got %d", v)
	}
	if !s.IsEmpty() {
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
	if !s.IsEmpty() {
		t.Errorf("Stack failed. Expected True, Got False")
	}
}
