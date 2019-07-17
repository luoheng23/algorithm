package chapter18

// BNode for BTree
type BNode struct {
	n int  // number of keys
	keys []int
	leaf bool // leaf or not
	c []*BNode // children (n+1)
}

// BTree is b tree
type BTree struct {
	root *BNode
	t    int
}

func diskRead(x *BNode, c *BNode) {}
func diskWrite(x *BNode) {}

// Create a BTree
func Create() BTree {
	return BTree{
		&BNode{leaf: true, n: 0}
	}
}

// Search for key
func (b *BTree) Search(x *BNode, k int) (*BNode, int){
	i := 0
	for i = 0; i < x.n && k > x.keys[i]; i++ {}
	if i < x.n && k == x.keys[i] {
		return x, i
	} else if x.leaf {
		return nil, 0
	}
	diskRead(x, x.c[i])
	return b.Search(x.c[i], k)
}

func (b *BTree) splitChild(x *BNode, i int) {
	y := x.c[i]
	z := &BNode{leaf: y.leaf, n: b.t-1}
	z.key = y.key[b.t:]
	if !y.leaf {
		z.c = y.c[b.t:]
	}
	y.n -= t - 1
	x.c = append(x.c, nil)
	x.c[i+2:] = x.c[i+1:x.n+1]
	x.c[i+1] = z
	x.keys = append(x.keys, nil)
	x.keys[i+1:] = x.keys[i:x.n]
	x.key[i] = y.key[i]
	x.n++
	diskWrite(y)
	diskWrite(z)
	diskWrite(x)
}

func (b *BTree) Insert(k int) {
	r := b.root
	if r.n == 2 * b.t - 1 {
		s := &BNode{leaf: false, n: 0}
		s.c = make([]int, 4)
		s.c[0] = r
		b.splitChild(s, 1)
		b.insertNotFull(s, k)
	} else {
		b.insertNotFull(r, k)
	}
}

func (b *BTree) insertNotFull(x *BNode, k int) {
	i := x.n
	if x.leaf {
		for i >= 0 && k < x.key[i] {
			x.keys[i+1] = x.keys[i]
			i--
		}
		x.keys[i+1] = k
		x.n++
		diskWrite(x)
	} else {
		for i >= 0 && k < x.key[i] {
			i--
		}
		i++
		diskRead(x, c[i])
		if x.c[i].n == 2*b.t-1 {
			b.splitChild(x, i)
			if k > x.key[i] {
				i++
			}
		}
		b.insertNotFull(x.c[i], k)
	}
}
