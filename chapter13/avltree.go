package chapter13

import "math"

type height int

// AVLNode of tree
type AVLNode struct {
	key         int
	h           height
	left, right *AVLNode
	p           *AVLNode
}

// AVLTree is AVL tree
type AVLTree struct {
	null *AVLNode
	root *AVLNode
}

// CreateAVLTree create a tree
func CreateAVLTree(A []int) AVLTree {
	r := AVLTree{
		null: &AVLNode{},
	}
	r.root = r.null
	for _, a := range A {
		r.Insert(&AVLNode{key: a})
	}
	return r
}

// InOrder performs inorder walk
func (r *AVLTree) InOrder(x *AVLNode) []int {
	var A []int
	if x != r.null {
		A = append(A, r.InOrder(x.left)...)
		A = append(A, x.key)
		A = append(A, r.InOrder(x.right)...)
	}
	return A
}

// PreOrder performs preorder walk
func (r *AVLTree) PreOrder(x *AVLNode) []int {
	var A []int
	if x != r.null {
		A = append(A, x.key)
		A = append(A, r.InOrder(x.left)...)
		A = append(A, r.InOrder(x.right)...)
	}
	return A
}

// PostOrder performs postorder walk
func (r *AVLTree) PostOrder(x *AVLNode) []int {
	var A []int
	if x != r.null {
		A = append(A, r.InOrder(x.left)...)
		A = append(A, r.InOrder(x.right)...)
		A = append(A, x.key)
	}
	return A
}

// Search for a key
func (r *AVLTree) Search(x *AVLNode, k int) *AVLNode {
	if x == r.null || k == x.key {
		return x
	} else if k < x.key {
		return r.Search(x.left, k)
	} else {
		return r.Search(x.right, k)
	}
}

// QuickSearch for a key
func (r *AVLTree) QuickSearch(x *AVLNode, k int) *AVLNode {
	for x != r.null && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// Min minimum AVLNode
func (r *AVLTree) Min(x *AVLNode) *AVLNode {
	for x.left != r.null {
		x = x.left
	}
	return x
}

// Max maximum AVLNode
func (r *AVLTree) Max(x *AVLNode) *AVLNode {
	for x.right != r.null {
		x = x.right
	}
	return x
}

// Successor previous AVLNode
func (r *AVLTree) Successor(x *AVLNode) *AVLNode {
	if x.right != r.null {
		return r.Min(x.right)
	}
	y := x.p
	for y != r.null && x == y.right {
		x, y = y, y.p
	}
	return y
}

// Preccessor post AVLNode
func (r *AVLTree) Preccessor(x *AVLNode) *AVLNode {
	if x.left != r.null {
		return r.Max(x.left)
	}
	y := x.p
	for y != r.null && x == y.left {
		x, y = y, y.p
	}
	return y
}

func (r *AVLTree) leftRotate(x *AVLNode) {
	y := x.right
	x.right = y.left
	if y.left != r.null {
		y.left.p = x
	}
	y.p = x.p
	if x.p == r.null {
		r.root = y
	} else if x == x.p.left {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.left = x
	x.p = y
	x.h = max(x.left.h, x.right.h) + 1
	y.h = max(y.left.h, y.right.h) + 1
}

func (r *AVLTree) rightRotate(x *AVLNode) {
	y := x.left
	x.left = y.right
	if y.right != r.null {
		y.left.p = x
	}
	y.p = x.p
	if x.p == r.null {
		r.root = y
	} else if x == x.p.left {
		x.p.left = y
	} else {
		x.p.right = y
	}
	y.right = x
	x.p = y
	x.h = max(x.left.h, x.right.h) + 1
	y.h = max(y.left.h, y.right.h) + 1
}

func max(T ...height) height {
	max := height(0)
	for _, t := range T {
		if t > max {
			max = t
		}
	}
	return max
}

func (r *AVLTree) balance(z *AVLNode) {
	for z != r.null {
		if z.left.h > z.right.h+1 {
			if z.left.right.h > z.left.left.h {
				r.leftRotate(z.left)
			}
			r.rightRotate(z)
			z = z.p
		} else if z.right.h > z.left.h+1 {
			if z.right.left.h > z.right.right.h {
				r.rightRotate(z.right)
			}
			r.leftRotate(z)
			z = z.p
		} else {
			z.h = max(z.left.h, z.right.h) + 1
		}
		z = z.p
	}
}

// Insert a AVLNode
func (r *AVLTree) Insert(z *AVLNode) {
	y, x := r.null, r.root
	for x != r.null {
		y = x
		if z.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	z.p = y
	if y == r.null {
		r.root = z
	} else if z.key < y.key {
		y.left = z
	} else {
		y.right = z
	}
	z.left = r.null
	z.right = r.null
	z.h = 1
	r.balance(z.p)
}

func (r *AVLTree) transplant(u, v *AVLNode) {
	if u.p == r.null {
		r.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

// Delete a AVLNode
func (r *AVLTree) Delete(z *AVLNode) {
	y, x := z, z
	if z.left == r.null {
		x = z.right
		r.transplant(z, z.right)
	} else if z.right == r.null {
		x = z.left
		r.transplant(z, z.left)
	} else {
		y = r.Min(z.right)
		x := y.right
		if y.p == z {
			x.p = y
		} else {
			r.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
		r.transplant(z, y)
		y.left = z.left
		y.left.p = y
	}
	r.balance(x.p)
}

func (s *AVLTree) isBalance(r *AVLNode) bool {
	if r == s.null {
		return true
	}
	if math.Abs(float64(r.left.h-r.right.h)) >= 2 {
		return false
	}
	return s.isBalance(r.left) && s.isBalance(r.right)
}
