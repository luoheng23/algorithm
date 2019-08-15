package chapter12

import "fmt"

// Node for SearchTree
type Node struct {
	key         int
	left, right *Node
	p           *Node
}

// SearchTree represents a tree
type SearchTree struct {
	root *Node
}

// CreateSearchTree init a tree
func CreateSearchTree(A []int) SearchTree {
	s := SearchTree{}
	for _, a := range A {
		s.Insert(&Node{key: a})
	}
	return s
}

// InOrder O(n)
func (s *SearchTree) InOrder(x *Node) []int {
	var A []int
	if x != nil {
		A = append(A, s.InOrder(x.left)...)
		A = append(A, x.key)
		A = append(A, s.InOrder(x.right)...)
	}
	return A
}

// PreOrder O(n)
func (s *SearchTree) PreOrder(x *Node) []int {
	var A []int
	if x != nil {
		A = append(A, x.key)
		A = append(A, s.InOrder(x.left)...)
		A = append(A, s.InOrder(x.right)...)
	}
	return A
}

// PostOrder O(n)
func (s *SearchTree) PostOrder(x *Node) []int {
	var A []int
	if x != nil {
		A = append(A, s.InOrder(x.left)...)
		A = append(A, s.InOrder(x.right)...)
		A = append(A, x.key)
	}
	return A
}

// Search rescursive search
func (s *SearchTree) Search(x *Node, k int) *Node {
	if x == nil || k == x.key {
		return x
	} else if k < x.key {
		return s.Search(x.left, k)
	}
	return s.Search(x.right, k)
}

// QuickSearch not rescursive search
func (s *SearchTree) QuickSearch(x *Node, k int) *Node {
	for x != nil && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// Min rescursive
func (s *SearchTree) Min(x *Node) *Node {
	if x.left == nil {
		return x
	}
	return s.Min(x.left)
}

// QuickMin minimum Node
func (s *SearchTree) QuickMin(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

// Max rescursive
func (s *SearchTree) Max(x *Node) *Node {
	if x.right == nil {
		return x
	}
	return s.Max(x.right)
}

// QuickMax maximum Node
func (s *SearchTree) QuickMax(x *Node) *Node {
	for x.right != nil {
		x = x.right
	}
	return x
}

// Successor previous Node O(h)
func (s *SearchTree) Successor(x *Node) *Node {
	if x.right != nil {
		return s.Min(x.right)
	}
	y := x.p
	for y != nil && x == y.right {
		x, y = y, y.p
	}
	return y
}

// Preccessor latter Node O(h)
func (s *SearchTree) Preccessor(x *Node) *Node {
	if x.left != nil {
		return s.Max(x.left)
	}
	y := x.p
	for y != nil && x == y.left {
		x, y = y, y.p
	}
	return y
}

// Insert O(lgn)
func (s *SearchTree) Insert(z *Node) {
	x := s.root
	y := x
	for x != nil {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == nil {
		s.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
}

func (s *SearchTree) transplant(u, v *Node) {
	if u.p == nil {
		s.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	if v != nil {
		v.p = u.p
	}
}

// Delete O(h)
func (s *SearchTree) Delete(z *Node) {
	if z.left == nil {
		s.transplant(z, z.right)
	} else if z.right == nil {
		s.transplant(z, z.left)
	} else {
		y := s.Min(z.right)
		if y.p != z {
			s.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		s.transplant(z, y)
		y.left = z.left
		y.left.p = y
	}
}

// maximum priority queue
// method:
// InsertMax
// Maximum
// ExtractMax

// InsertMax O(h)
func (s *SearchTree) InsertMax(x int) {
	s.Insert(&Node{key: x})
}

// Maximum O(h)
func (s *SearchTree) Maximum() int {
	max := s.Max(s.root)
	if max != nil {
		return max.key
	}
	return -1
}

// ExtractMax O(h)
func (s *SearchTree) ExtractMax() (int, error) {
	if s.root == nil {
		return 0, fmt.Errorf("Heap overflow")
	}
	max := s.Max(s.root)
	s.Delete(max)
	return max.key, nil
}

// minimum priority queue
// method:
// Minimum
// ExtractMin
// InsertMin

// InsertMin O(h)
func (s *SearchTree) InsertMin(x int) {
	s.Insert(&Node{key: x})
}

// Minimum O(h)
func (s *SearchTree) Minimum() int {
	min := s.Min(s.root)
	if min != nil {
		return min.key
	}
	return -1
}

// ExtractMin O(h)
func (s *SearchTree) ExtractMin() (int, error) {
	if s.root == nil {
		return 0, fmt.Errorf("Heap overflow")
	}
	min := s.Min(s.root)
	s.Delete(min)
	return min.key, nil
}
