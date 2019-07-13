package chapter22

import "fmt"

const (
	white = iota
	gray
	black
)

// Bfs is bfs for graph
func Bfs(graph Graph, s Vertex) {
	for i, v := range graph.table {
		if v != s {
			graph.table[i].key = white
		}
	}
	s.key = gray
	Q := make([]Vertex, len(graph.table))
	Q[0] = s
	cur := 1
	for cur != 0 {
		cur--
		u := Q[cur]
		n := u.next
		for n != nil {
			if graph.table[n.v].key == white {
				graph.table[n.v].key = gray
				Q[cur] = graph.table[n.v]
				cur++
			}
			n = n.next
		}
		fmt.Println(u.v)
		graph.table[u.v].key = black
	}
}