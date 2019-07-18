package chapter21

// Node used in findSet
type Node struct {
	rank int
	p *Node
	key rune
}

// MakeSet make x a Node as a set
func MakeSet(x *Node) {
	x.rank = 0
	x.p = x
}

// Union different sets
func Union(x, y *Node) {
	link(FindSet(x), FindSet(y))
}

func link(x, y *Node) {
	if x.rank > y.rank {
		y.p = x
	} else {
		x.p = y
		if x.rank == y.rank {
			y.rank++
		}
	}
}

// FindSet return the set belong
func FindSet(x *Node) *Node{
	if x != x.p {
		x.p = FindSet(x.p)
	}
	return x.p
}

