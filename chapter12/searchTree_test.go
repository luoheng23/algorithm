package chapter12

import "testing"

func TestSearchTree(t *testing.T) {
	A := [10]int{5, 3, 9, 0, 4, 8, 6, 2, 7, 1}
	sortedA := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := CreateSearchTree(A[:])
	for i, a := range s.InOrder(s.root) {
		if a != sortedA[i] {
			t.Errorf("SearchTree: Create failed. Expected %d, Got %d", sortedA[i], a)
		}
	}
	if s.Max(s.root).key != 9 || s.Min(s.root).key != 0 {
		t.Errorf("SearchTree: Min and Max failed. Expected (0, 9), Got (%d, %d)", s.Min(s.root).key, s.Max(s.root).key)
	}
	if s.Successor(s.Search(s.root, 4)).key != 5 {
		t.Errorf("SearchTree: Successor failed. Expected 5, Got %d", s.Successor(&Node{key: 4}).key)
	}
	if s.Preccessor(s.Search(s.root, 5)).key != 4 {
		t.Errorf("SearchTree: Preccessor failed. Expected 4, Got %d", s.Preccessor(s.Search(s.root, 5)).key)
	}
	s.Delete(s.Search(s.root, 5))
	if s.root.key != 6 {
		t.Errorf("SearchTree: Delete failed. Expected 6, Got %d", s.root.key)
	}

}
