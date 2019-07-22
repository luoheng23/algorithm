
from graph import initGraph

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

def kruskal(G, w):
    A = []
    for v in G.V:
        makeSet(G.V[v])
    G.E.sort()
    for u, v in G.E:
        if findSet(G.V[v]) != findSet(G.V[u]):
            A.append((u, v))
            union(G.V[u], G.V[v])
    return A

