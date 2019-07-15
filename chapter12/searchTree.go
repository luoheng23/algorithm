package chapter12

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
	} else {
		return s.Search(x.right, k)
	}
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

// Min minimum Node
func (s *SearchTree) Min(x *Node) *Node {
	for x.left != nil {
		x = x.left
	}
	return x
}

// Max maximum Node
func (s *SearchTree) Max(x *Node) *Node {
	for x.right != nil {
		x = x.right
	}
	return x
}

// Successor previous Node
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

// Preccessor latter Node
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

// Delete O(1)
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

