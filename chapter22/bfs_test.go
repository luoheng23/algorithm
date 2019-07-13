package chapter22

import "testing"

// TestBfs
func TestBfs(t *testing.T) {
	vs := []vertex {
		vertex{1, 5},
		vertex{2, 6},
		vertex{3, 7},
		vertex{4, 8},
		vertex{5, 9},
		vertex{6, 10},
		vertex{7, 11},
	}
	es := []edge {
		edge{1, 2, 3},
		edge{2, 4, 4},
		edge{1, 7, 6},
		edge{1, 5, 2},
		edge{2, 7, 4},
		edge{7, 3, 3},
		edge{6, 1, 9},
		edge{5, 2, 4},
		edge{4, 2, 1},
	}
	graph := InitGraph(vs, es, 7)
	Bfs(graph, graph.table[1])
}