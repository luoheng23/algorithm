
import sys
sys.path.insert(0, '..')

from chapter22.graph import initGraph
from relax import init, relax

def bellmanFord(G, s):
    init(G, s)
    for _ in range(len(G.V)-1):
        for u, v, _ in G.E:
            relax(G, u, v)
    for u, v, _ in G.E:
        if G.V[v].d > G.V[u].d + G.W[u, v]:
            return False
    return True

def main():
    vertex = ['s', 't', 'y', 'x', 'z']
    edges = [('s', 't', 6), ('s', 'y', 7), ('t', 'x', 5), ('x', 't', -2),
            ('t', 'z', -4), ('y', 'z', 9), ('y', 'x', -3), ('z', 'x', 7),
            ('z', 's', 2), ('t', 'y', 8)]
    G = initGraph(vertex, edges)
    bellmanFord(G, 's')
    for v in G.V:
        print(G.V[v].d)

if __name__ == "__main__":
    main()