
from graph import initGraph

time = 0

def dfs(G):
    global time
    time = 0
    for v in G.V:
        if G.V[v].color == G.WHITE:
            dfsVisit(G, v)

def dfsVisit(G, v):
    global time
    time += 1
    G.V[v].d = time
    G.V[v].color = G.GRAY
    for u in G.Adj[v]:
        if G.V[u].color == G.WHITE:
            G.V[u].pi = v
            dfsVisit(G, u)
    G.V[v].color = G.BLACK
    time += 1
    G.V[v].f = time

def main():
    vertex = [0, 1, 2, 3, 4, 5, 6, 7]
    edges = [(0, 1), (1, 2), (2, 3), (3, 4), (3, 5),
             (4, 5), (4, 7), (4, 6), (5, 6), (6, 7)]
    G = initGraph(vertex, edges)
    dfs(G)
    for v in G.V:
        print(v, G.V[v].d, G.V[v].f)

if __name__ == "__main__":
    main()

