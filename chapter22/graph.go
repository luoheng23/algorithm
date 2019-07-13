package chapter22

// use to construct graph
type edge struct {
	u int
	v int
	weight int
}

type vertex struct {
	v int
	key int
}

// Vertex represent vertex
type Vertex struct {
	key int
	next *Node
}

// Node represent node
type Node struct {
	v int
	next *Node
}

// Graph represent graph
type Graph struct {
	table []Vertex
	len int
}

type Matrix struct {
	vs []int
	matrix [][]int
	len int
}

// InitGraph init a graph, the vertex should be 1 to n
func InitGraph(vs []vertex, es []edge, n int) Graph {
	graph := Graph{len: n+1}
	graph.table = make([]Vertex, n+1)
	for _, v := range vs {
		graph.table[v.v].key = v.key
	}
	for _, e := range es {
		node := Node{v: e.v}
		node.next = graph.table[e.u].next
		graph.table[e.u].next = &node
	}
	return graph
}

func InitMatrix(vs []vertex, es []edge, n int) Matrix {
	m := Matrix{len: n}
	m.vs = make([]int, n+1)
	for _, v := range vs {
		m.vs[v.v] = v.key
	}
	m.matrix = make([][]int, n+1)
	for i:= 0; i < n + 1; i++ {
		m.matrix[i] = make([]int, n+1)
	}
	for _, e := range es {
		m.matrix[e.u][e.v] = e.weight
	}
	return m
}