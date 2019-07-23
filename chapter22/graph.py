
from collections import defaultdict


class Vertex:

    def __init__(self, value):
        self.value = value

    def __str__(self):
        return f"{self.value}"

    def __repr__(self):
        return self.__str__()
    
    def __eq__(self, value):
        return self.value == value


class Graph:
    def __init__(self, vertex, edges):
        self.V = {num: Vertex(num) for num in vertex}
        self.Adj = defaultdict(list)
        self.E = edges
        self.W = {}
        for edge in edges:
            u, v, *w = edge
            self.Adj[u].append(v)
            self.Adj[v].append(u)
            if w != []:
                self.W[u, v] = w[0]
                self.W[v, u] = w[0]

    def __str__(self):
        return f"{self.V}, {self.E}"


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
