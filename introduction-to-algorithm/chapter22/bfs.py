
from graph import initGraph

WHITE = 0
GRAY = 1
BLACK = 2

def bfs(G, s):
    for v in G.V:
        G.V[v].color = WHITE
        G.V[v].d = float("inf")
        G.V[v].pi = None
    G.V[s].color, G.V[s].d, G.V[s].pi = GRAY, 0, None
    q = [s]
    while q != []:
        u = q.pop()
        for v in G.Adj[u]:
            if G.V[v].color == WHITE:
                G.V[v].color = GRAY
                G.V[v].d += 1
                G.V[v].pi = u
                q.insert(0, v)
        G.V[u].color = BLACK


def printPath(G, s, v):
    if v == s:
        print(s)
    elif G.V[v].pi is None:
        print(f"no path from {s} to {v} exists.")
    else:
        printPath(G, s, G.V[v].pi)
        print(v)


def main():
    vertex = [0, 1, 2, 3, 4, 5, 6, 7]
    edges = [(0, 1), (1, 2), (2, 3), (3, 4), (3, 5),
             (4, 5), (4, 7), (4, 6), (5, 6), (6, 7)]
    G = initGraph(vertex, edges)
    bfs(G, 3)
    printPath(G, 3, 7)


if __name__ == "__main__":
    main()
