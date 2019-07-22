
from graph import initGraph


def bfs(G, s):
    G.V[s].color, G.V[s].d, G.V[s].pi = G.GRAY, 0, None
    q = [s]
    while q != []:
        u = q.pop()
        print(u)
        for v in G.Adj[u]:
            if G.V[v].color == G.WHITE:
                G.V[v].color = G.GRAY
                G.V[v].d += 1
                G.V[v].pi = u
                q.insert(0, v)
        G.V[u].color = G.BLACK


def main():
    vertex = [0, 1, 2, 3, 4, 5, 6, 7]
    edges = [(0, 1), (1, 2), (2, 3), (3, 4), (3, 5),
             (4, 5), (4, 7), (4, 6), (5, 6), (6, 7)]
    G = initGraph(vertex, edges)
    bfs(G, 3)


if __name__ == "__main__":
    main()
