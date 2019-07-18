package chapter18

// BNode for BTree
type BNode struct {
	n    int // number of keys
	keys []rune
	leaf bool     // leaf or not
	c    []*BNode // children (n+1)
}

// BTree is b tree
type BTree struct {
	root *BNode
	t    int
}

func diskRead(x *BNode, c *BNode) {}
func diskWrite(x *BNode)          {}

// CreateNode return a node
func (b *BTree) CreateNode() *BNode {
	return &BNode{
		n:    0,
		keys: make([]rune, 2*b.t-1),
		leaf: true,
		c:    make([]*BNode, 2*b.t),
	}
}

// CreateTree returns a BTree
func CreateTree(t int) BTree {
	b := BTree{}
	b.t = t
	b.root = b.CreateNode()
	return b
}

// Search for keys
func (b *BTree) Search(x *BNode, k rune) (*BNode, rune) {
	i := 0
	for i = 0; i < x.n && k > x.keys[i]; i++ {
	}
	if i < x.n && k == x.keys[i] {
		return x, x.keys[i]
	} else if x.leaf {
		return nil, 0
	}
	diskRead(x, x.c[i])
	return b.Search(x.c[i], k)
}

func (b *BTree) splitChild(x *BNode, i int) {
	y := x.c[i]
	z := b.CreateNode()
	copy(z.keys, y.keys[b.t:y.n])
	if !y.leaf {
		copy(z.c, y.c[b.t:y.n+1])
	}
	z.n = b.t - 1
	y.n = b.t - 1
	x.n++
	copy(x.c[i+2:x.n+1], x.c[i+1:x.n])
	x.c[i+1] = z
	copy(x.keys[i+1:x.n], x.keys[i:x.n-1])
	x.keys[i] = y.keys[y.n]
	diskWrite(y)
	diskWrite(z)
	diskWrite(x)
}

// Insert a node
func (b *BTree) Insert(k rune) {
	r := b.root
	if r.n == 2*b.t-1 {
		s := b.CreateNode()
		b.root = s
		s.c[0] = r
		s.leaf = false
		b.splitChild(s, 0)
		b.insertNotFull(s, k)
	} else {
		b.insertNotFull(r, k)
	}
}

func (b *BTree) insertNotFull(x *BNode, k rune) {
	i := x.n - 1
	if x.leaf {
		for i >= 0 && k < x.keys[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
		x.n++
		diskWrite(x)
	} else {
		for i >= 0 && k < x.keys[i] {
			i--
		}
		i++
		diskRead(x, x.c[i])
		if x.c[i].n == 2*b.t-1 {
			b.splitChild(x, i)
			if k > x.keys[i] {
				i++
			}
		}
		b.insertNotFull(x.c[i], k)
	}
}

// Delete a node
func (b *BTree) Delete(x *BNode, k rune){
	i := 0
	for ; i < x.n; i++ {
		if x.keys[i] >= k {
			break
		}
	}
	if i < x.n && x.keys[i] == k {
		if x.leaf {
			copy(x.keys[i:], x.keys[i+1:])
			x.n--
		} else if x.c[i].n >= b.t {
			x.keys[i] = x.c[i].keys[x.c[i].n-1]
			b.Delete(x.c[i], x.keys[i])
		} else if x.c[i+1].n >= b.t {
			x.keys[i] = x.c[i+1].keys[0]
			b.Delete(x.c[i+1], x.keys[i])
		} else {
			x.c[i].keys[x.c[i].n] = x.keys[i]
			x.c[i].n++
			copy(x.c[i].keys[x.c[i].n:], x.c[i+1].keys[:x.c[i+1].n])
			x.c[i].n += x.c[i+1].n
			copy(x.c[i].c[x.n+1:], x.c[i+1].c[:x.c[i+1].n+1])
			copy(x.keys[i:], x.keys[i+1:])
			copy(x.c[i+1:], x.c[i+2:])
			x.n--
			b.Delete(x.c[i], k)
		}
	} else if i < x.n && x.keys[i] > k && !x.leaf {
		if x.c[i].n < b.t {
			x.c[i].keys[x.c[i].n] = x.keys[i]
			x.c[i].n++
			if x.c[i+1].n >= b.t {
				x.keys[i] = x.c[i+1].keys[0]
				x.c[i].c[x.c[i].n++] = x.c[i+1].c[0] 
				b.Delete(x.c[i+1], x.keys[i])
				b.Delete(x.c[i], k)
			} else {
				copy(x.c[i].keys[x.c[i].n:], x.c[i+1].keys[:x.c[i+1].n])
				x.c[i].n += x.c[i+1].n
				copy(x.c[i].c[x.n+1:], x.c[i+1].c[:x.c[i+1].n+1])
				copy(x.keys[i:], x.keys[i+1:])
				copy(x.c[i+1:], x.c[i+2:])
				x.n--
			}
		}
		b.Delete(x.c[i], k)
	}
	if x.n == 0 && x == b.root {
		b.root = x.c[0]
	}
}

func (b *BTree) inOrder(s *BNode) []rune {
	A := make([]rune, 0, 4)
	if s != nil {
		for i, key := range s.keys {
			if i < s.n {
				A = append(A, b.inOrder(s.c[i])...)
				A = append(A, key)
			}
		}
		A = append(A, b.inOrder(s.c[s.n])...)
	}
	return A
}
