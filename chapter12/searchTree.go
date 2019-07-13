package chapter12

type node struct {
	key         int
	left, right *node
	p           *node
}

type SearchTree struct {
	root *node
}

func CreateSearchTree(A []int) SearchTree {
	s := SearchTree{}
	for _, a := range A {
		s.Insert(&node{key: a})
	}
	return s
}

func (s *SearchTree) InOrder(x *node) []int {
	var A []int
	if x != nil {
		A = append(A, s.InOrder(x.left)...)
		A = append(A, x.key)
		A = append(A, s.InOrder(x.right)...)
	}
	return A
}

func (s *SearchTree) PreOrder(x *node) []int {
	var A []int
	if x != nil {
		A = append(A, x.key)
		A = append(A, s.InOrder(x.left)...)
		A = append(A, s.InOrder(x.right)...)
	}
	return A
}

func (s *SearchTree) PostOrder(x *node) []int {
	var A []int
	if x != nil {
		A = append(A, s.InOrder(x.left)...)
		A = append(A, s.InOrder(x.right)...)
		A = append(A, x.key)
	}
	return A
}

func (s *SearchTree) Search(x *node, k int) *node {
	if x == nil || k == x.key {
		return x
	} else if k < x.key {
		return s.Search(x.left, k)
	} else {
		return s.Search(x.right, k)
	}
}

func (s *SearchTree) QuickSearch(x *node, k int) *node {
	for x != nil && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

func (s *SearchTree) Min(x *node) *node {
	for x.left != nil {
		x = x.left
	}
	return x
}

func (s *SearchTree) Max(x *node) *node {
	for x.right != nil {
		x = x.right
	}
	return x
}

// InOrder
func (s *SearchTree) Successor(x *node) *node {
	if x.right != nil {
		return s.Min(x.right)
	}
	y := x.p
	for y != nil && x == y.right {
		x, y = y, y.p
	}
	return y
}

// InOrder
func (s *SearchTree) Preccessor(x *node) *node {
	if x.left != nil {
		return s.Max(x.left)
	}
	y := x.p
	for y != nil && x == y.left {
		x, y = y, y.p
	}
	return y
}

func (s *SearchTree) Insert(z *node) {
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

func (s *SearchTree) transplant(u, v *node) {
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

func (s *SearchTree) Delete(z *node) {
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
