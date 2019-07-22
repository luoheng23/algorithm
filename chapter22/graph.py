
from collections import defaultdict


class Vertex(int):

    def __init__(self, value):
        self.color = Graph.WHITE
        self.d = float("inf")
        self.pi = None


class Graph:
    WHITE = 0
    GRAY = 1
    BLACK = 2

    def __init__(self, vertex, edges):
        self.V = {num: Vertex(num) for num in vertex}
        self.Adj = defaultdict(list)
        self.E = edges
        for u, v in edges:
            self.Adj[u].append(v)

    def __str__(self):
        return f"{self.V}, {self.Adj}"


def initGraph(vertex, edges):
    return Graph(vertex, edges)


def main():
    vertex = [0, 1, 2, 3, 4, 5, 6, 7]
    edges = [(0, 1), (1, 2), (2, 3), (3, 4), (3, 5),
             (4, 5), (4, 7), (4, 6), (5, 6), (6, 7)]
    graph = initGraph(vertex, edges)
    print(graph)


if __name__ == "__main__":
    main()
