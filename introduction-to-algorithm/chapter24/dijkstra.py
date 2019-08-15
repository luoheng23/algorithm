
import sys
sys.path.insert(0, '..')

from chapter22.graph import initGraph
from chapter23.prim import MinHeap
from relax import init, relax

def dijkstra(G, s):
    init(G, s)
    for v in G.V:
        G.V[v].key = float("inf")
    S = []
    heap = MinHeap(list(G.V.values()))
    while not heap.isEmpty():
        u = heap.extractMin()
        S.append(u.value)
        for v in G.Adj[u.value]:
            relax(G, u.value, v)


def main():
    vertex = ['s', 't', 'y', 'x', 'z']
    edges = [('s', 't', 10), ('s', 'y', 5), ('t', 'x', 1),
            ('y', 'z', 2), ('y', 'x', 9), ('z', 'x', 6),
            ('z', 's',72), ('t', 'y', 2), ('t', 'y', 2), ('y', 't', 3)]
    G = initGraph(vertex, edges)
    dijkstra(G, 's')
    for v in G.V:
        print(G.V[v].d)

if __name__ == "__main__":
    main()