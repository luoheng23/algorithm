package chapter21

import "testing"

type edge struct {
	i rune
	j rune
}

func TestSet(t *testing.T) {
	v := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}	
	e := []edge{
		edge{'a', 'c'},
		edge{'a', 'b'},
		edge{'b', 'c'},
		edge{'b', 'd'},
		edge{'e', 'g'},
		edge{'e', 'f'},
		edge{'h', 'i'},
	}
	V := make(map[rune]*Node, 10)
	for _, r := range v {
		V[r] = &Node{
			key: r,
		}
		MakeSet(V[r])
	}
	for _, z := range e {
		if FindSet(V[z.i]) != FindSet(V[z.j]) {
			Union(V[z.i], V[z.j])
		}
	}
	for _, z := range e {
		if FindSet(V[z.i]) != FindSet(V[z.j]) {
			t.Errorf("Set failed. Expected (%c, %c) true, Got false", z.i, z.j)
		}
	}
}