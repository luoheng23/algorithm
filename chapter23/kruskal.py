
import sys
sys.path.insert(0, '..')

from chapter22.graph import initGraph


def makeSet(v):
    v.rank = 0
    v.p = v


def union(u, v):
    link(findSet(u), findSet(v))


def link(u, v):
    if u.rank > v.rank:
        v.p = u
    else:
        if u.rank == v.rank:
            v.rank += 1
        u.p = v


def findSet(u):
    if u.p != u:
        u.p = findSet(u.p)
    return u.p


def kruskal(G):
    A = []
    for v in G.V:
        makeSet(G.V[v])
    G.E.sort(key=lambda x: x[2])
    for u, v, _ in G.E:
        if findSet(G.V[v]) != findSet(G.V[u]):
            A.append((u, v))
            union(G.V[u], G.V[v])
    return A


def main():
    vertex = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i']
    edges = [('a', 'b', 4), ('a', 'h', 8), ('b', 'c', 7),  ('c', 'd', 7), ('d', 'e', 9),
             ('e', 'f', 10), ('f', 'g', 2), ('g',
                                             'h', 1), ('h', 'i', 7), ('g', 'i', 6),
             ('i', 'c', 2), ('f', 'c', 4)]
    graph = initGraph(vertex, edges)
    print(kruskal(graph))


if __name__ == "__main__":
    main()
