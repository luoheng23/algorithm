// Package chapter13 is RBTree
// RBTree definition:
// 1. all nodes are red or black
// 2. root is black
// 3. leaf is black
// 4. if Node is red, its' children are black
// 5. for each Node, from it to leaf, all include same number of black Node
package chapter13

type color bool

const (
	black color = true
	red   color = false
)

// Node of tree
type Node struct {
	key         int
	color       color
	left, right *Node
	p           *Node
}

// RBTree is black and red tree
type RBTree struct {
	null *Node
	root *Node
}

// CreateRBTree create a tree
func CreateRBTree(A []int) RBTree {
	r := RBTree{
		null: &Node{color: black},
	}
	r.root = r.null
	for _, a := range A {
		r.Insert(&Node{key: a})
	}
	return r
}

// InOrder performs inorder walk
func (r *RBTree) InOrder(x *Node) []int {
	var A []int
	if x != r.null {
		A = append(A, r.InOrder(x.left)...)
		A = append(A, x.key)
		A = append(A, r.InOrder(x.right)...)
	}
	return A
}

// PreOrder performs preorder walk
func (r *RBTree) PreOrder(x *Node) []int {
	var A []int
	if x != r.null {
		A = append(A, x.key)
		A = append(A, r.InOrder(x.left)...)
		A = append(A, r.InOrder(x.right)...)
	}
	return A
}

// PostOrder performs postorder walk
func (r *RBTree) PostOrder(x *Node) []int {
	var A []int
	if x != r.null {
		A = append(A, r.InOrder(x.left)...)
		A = append(A, r.InOrder(x.right)...)
		A = append(A, x.key)
	}
	return A
}

// Search for a key
func (r *RBTree) Search(x *Node, k int) *Node {
	if x == r.null || k == x.key {
		return x
	} else if k < x.key {
		return r.Search(x.left, k)
	} else {
		return r.Search(x.right, k)
	}
}

// QuickSearch for a key
func (r *RBTree) QuickSearch(x *Node, k int) *Node {
	for x != r.null && k != x.key {
		if k < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	return x
}

// Min minimum Node
func (r *RBTree) Min(x *Node) *Node {
	for x.left != r.null {
		x = x.left
	}
	return x
}

// Max maximum Node
func (r *RBTree) Max(x *Node) *Node {
	for x.right != r.null {
		x = x.right
	}
	return x
}

// Successor previous Node
func (r *RBTree) Successor(x *Node) *Node {
	if x.right != r.null {
		return r.Min(x.right)
	}
	y := x.p
	for y != r.null && x == y.right {
		x, y = y, y.p
	}
	return y
}

// Preccessor post Node
func (r *RBTree) Preccessor(x *Node) *Node {
	if x.left != r.null {
		return r.Max(x.left)
	}
	y := x.p
	for y != r.null && x == y.left {
		x, y = y, y.p
	}
	return y
}

func (r *RBTree) leftRotate(x *Node) {
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
}

func (r *RBTree) rightRotate(x *Node) {
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
}

func (r *RBTree) insertFixUp(z *Node) {
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

// Insert a Node
func (r *RBTree) Insert(z *Node) {
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
	r.insertFixUp(z)
}

func (r *RBTree) transplant(u, v *Node) {
	if u.p == r.null {
		r.root = v
	} else if u == u.p.left {
		u.p.left = v
	} else {
		u.p.right = v
	}
	v.p = u.p
}

func (r *RBTree) deleteFixUp(x *Node) {
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

// Delete a Node
func (r *RBTree) Delete(z *Node) {
	y, yColor := z, z.color
	if z.left == r.null {
		z = z.right
		r.transplant(z, z.right)
	} else if z.right == r.null {
		z = z.left
		r.transplant(z, z.left)
	} else {
		y = r.Min(z.right)
		yColor = y.color
		x := y.right
		if y.p == z {
			x.p = y
		} else {
			r.transplant(y, y.right)
			y.right = z.right
			y.right.p = y
		}
	}
	r.transplant(z, y)
	y.left = z.left
	y.left.p = y
	y.color = z.color
	if yColor == black {
		r.deleteFixUp(y)
	}
}
