
import sys
sys.path.insert(0, '..')

from chapter22.graph import initGraph
from relax import init, relax

WHITE = 0
GRAY = 1
BLACK = 2

time = 0
L = []

def topoSort(G):
    global time, L
    time = 0
    L = []
    for v in G.V:
        G.V[v].color = WHITE
    for v in G.V:
        if G.V[v].color == WHITE:
            dfsVisit(G, v)

def dfsVisit(G, v):
    global time
    time += 1
    G.V[v].d = time
    G.V[v].color = GRAY
    for u in G.Adj[v]:
        if G.V[u].color == WHITE:
            G.V[u].pi = v
            dfsVisit(G, u)
    G.V[v].color = BLACK
    time += 1
    G.V[v].f = time
    L.insert(0, v)

def noCircle(G, s):
    topoSort(G)
    init(G, s)
    for u in L:
        for v in G.Adj[u]:
            relax(G, u, v)

def main():
    vertex = ['r', 's', 't', 'y', 'x', 'z']
    edges = [('r', 's', 5), ('r', 't', 3), ('s', 'x', 6), ('s', 't', 2),
            ('t', 'y', 4), ('t', 'z', 2), ('t', 'x', 7), ('x', 'y', -1),
            ('x', 'z', 1), ('y', 'z', -2)]
    G = initGraph(vertex, edges)
    noCircle(G, 's')
    for v in G.V:
        print(G.V[v], G.V[v].d)

if __name__ == "__main__":
    main()