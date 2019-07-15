// Package chapter14 is StatTree
// add size to RBTree
package chapter14

type color bool

const (
	black color = true
	red   color = false
)

// StatNode of tree
type StatNode struct {
	key         int
	color       color
	size        int
	left, right *StatNode
	p           *StatNode
}

// StatTree is black and red tree
type StatTree struct {
	null *StatNode
	root *StatNode
}

// CreateStatTree create a tree
func CreateStatTree(A []int) StatTree {
	r := StatTree{
		null: &StatNode{color: black, size: 0},
	}
	r.root = r.null
	for _, a := range A {
		r.Insert(&StatNode{key: a})
	}
	return r
}

// InOrder performs inorder walk
func (r *StatTree) InOrder(x *StatNode) []int {
	var A []int
	if x != r.null {
		A = append(A, r.InOrder(x.left)...)
		A = append(A, x.key)
		A = append(A, r.InOrder(x.right)...)
	}
	return A
}

// PreOrder performs preorder walk
func (r *StatTree) PreOrder(x *StatNode) []int {
	var A []int
	if x != r.null {
		A = append(A, x.key)
		A = append(A, r.InOrder(x.left)...)
		A = append(A, r.InOrder(x.right)...)
	}
	return A
}

// PostOrder performs postorder walk
func (r *StatTree) PostOrder(x *StatNode) []int {
	var A []int
	if x != r.null {
		A = append(A, r.InOrder(x.left)...)
		A = append(A, r.InOrder(x.right)...)
		A = append(A, x.key)
	}
	return A
}

// Search for a key
func (r *StatTree) Search(x *StatNode, k int) *StatNode {
	if x == r.null || k == x.key {
		return x
	} else if k < x.key {
		return r.Search(x.left, k)
	} else {
		return r.Search(x.right, k)
	}
}

// QuickSearch for a key
func (r *StatTree) QuickSearch(x *StatNode, k int) *StatNode {
	for x != r.null && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// Min minimum StatNode
func (r *StatTree) Min(x *StatNode) *StatNode {
	for x.left != r.null {
		x = x.left
	}
	return x
}

// Max maximum StatNode
func (r *StatTree) Max(x *StatNode) *StatNode {
	for x.right != r.null {
		x = x.right
	}
	return x
}

// Successor previous StatNode
func (r *StatTree) Successor(x *StatNode) *StatNode {
	if x.right != r.null {
		return r.Min(x.right)
	}
	y := x.p
	for y != r.null && x == y.right {
		x, y = y, y.p
	}
	return y
}

// Preccessor post StatNode
func (r *StatTree) Preccessor(x *StatNode) *StatNode {
	if x.left != r.null {
		return r.Max(x.left)
	}
	y := x.p
	for y != r.null && x == y.left {
		x, y = y, y.p
	}
	return y
}

func (r *StatTree) leftRotate(x *StatNode) {
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
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
}

func (r *StatTree) rightRotate(x *StatNode) {
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
	y.size = x.size
	x.size = x.left.size + x.right.size + 1
}

func (r *StatTree) insertFixUp(z *StatNode) {
	for z.p.color == red {
		if z.p == z.p.p.left {
			y := z.p.p.right
			if y.color == red {
				z.p.color, y.color, z.p.p.color = black, black, red
				z = z.p.p
			} else {
				if z == z.p.right {
					z = z.p
					r.leftRotate(z)
				}
				z.p.color, z.p.p.color = black, red
				r.rightRotate(z.p.p)
			}
		} else {
			y := z.p.p.left
			if y.color == red {
				z.p.color, y.color, z.p.p.color = black, black, red
				z = z.p.p
			} else {
				if z == z.p.left {
					z = z.p
					r.rightRotate(z)
				}
				z.p.color, z.p.p.color = black, red
				r.leftRotate(z.p.p)
			}
		}
	}
	r.root.color = black
}

// Insert a StatNode
func (r *StatTree) Insert(z *StatNode) {
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
	z.color = red
	z.size = 1
	x = z.p
	for x != r.null {
		x.size++
		x = x.p
	}
	r.insertFixUp(z)
}

func (r *StatTree) transplant(u, v *StatNode) {
	if u.p == r.null {
		r.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (r *StatTree) deleteFixUp(x *StatNode) {
	for x != r.root && x.color == black {
		if x == x.p.left {
			w := x.p.right
			if w.color == red {
				w.color = black
				x.p.color = red
				r.leftRotate(x.p)
				w = x.p.right
			}
			if w.left.color == black && w.right.color == black {
				w.color = red
				x = x.p
			} else if w.right.color == black {
				w.left.color = black
				w.color = red
				r.rightRotate(w)
				w = x.p.right
			}
			w.color = x.p.color
			x.p.color = black
			w.right.color = black
			r.leftRotate(x.p)
			x = r.root
		} else {
			w := x.p.left
			if w.color == red {
				w.color = black
				x.p.color = red
				r.rightRotate(x.p)
				w = x.p.left
			}
			if w.right.color == black && w.left.color == black {
				w.color = red
				x = x.p
			} else if w.left.color == black {
				w.right.color = black
				w.color = red
				r.leftRotate(w)
				w = x.p.left
			}
			w.color = x.p.color
			x.p.color = black
			w.left.color = black
			r.rightRotate(x.p)
			x = r.root
		}
	}
	x.color = black
}

// Delete a StatNode
func (r *StatTree) Delete(z *StatNode) {
	y, yColor := z, z.color
	x := z
	if z.left == r.null {
		x = z.right
		r.transplant(z, z.right)
	} else if z.right == r.null {
		x = z.left
		r.transplant(z, z.left)
	} else {
		y = r.Min(z.right)
		yColor = y.color
		x = y.right
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
		y.color = z.color
		y.size = z.size
	}
	t := x.p
	for t != r.null {
		t.size--
		t = t.p
	}
	if yColor == black {
		r.deleteFixUp(x)
	}
}

// Select the StatNode which is ith node
func (r *StatTree) Select(x *StatNode, i int) *StatNode {
	s := x.left.size + 1
	if i == s {
		return x
	} else if i < s {
		return r.Select(x.left, i)
	} else {
		return r.Select(x.right, i-s)
	}
}

// Rank the x node
func (r *StatTree) Rank(x *StatNode) int {
	s := x.left.size + 1
	y := x
	for y != r.root {
		if y == y.p.right {
			s += y.p.left.size + 1
		}
		y = y.p
	}
	return s
}

func (r *StatTree) isSizeRight(s *StatNode) bool {
	if s == r.null {
		return true
	}
	if s.left.size + s.right.size + 1 != s.size {
		return false
	}
	return r.isSizeRight(s.left) && r.isSizeRight(s.right)
}
