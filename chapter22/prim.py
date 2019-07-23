
from graph import initGraph


def parent(i):
    return i // 2


def left(i):
    return i*2


def right(i):
    return 2*i+1


class MinHeap:
    def __init__(self, List):
        self.minHeap = List
        self.heapSize = len(self.minHeap)
        self.buildMinHeap()

    def buildMinHeap(self):
        for i in range(self.heapSize//2+1, -1, -1):
            self.minHeapify(i)

    def minHeapify(self, i):
        minI = i
        if left(i) < self.heapSize and self.minHeap[left(i)].key <= self.minHeap[i].key:
            minI = left(i)
        if right(i) < self.heapSize and self.minHeap[right(i)].key <= self.minHeap[minI].key:
            minI = right(i)
        self.minHeap[minI], self.minHeap[i] = self.minHeap[i], self.minHeap[minI]
        if minI != i:
            self.minHeapify(minI)

    def Min(self):
        return self.minHeap[0]

    def extractMin(self):
        self.heapSize -= 1
        self.minHeap[0], self.minHeap[self.heapSize] = self.minHeap[self.heapSize], self.minHeap[0]
        self.minHeapify(0)
        return self.minHeap[self.heapSize]

    def isEmpty(self):
        return self.heapSize == 0

    def decreaseKey(self, i, key):
        self.minHeap[i].key = key
        while parent(i) >= 0 and self.minHeap[parent(i)].key > key:
            self.minHeap[i], self.minHeap[parent(
                i)] = self.minHeap[parent(i)], self.minHeap[i]
            i = parent(i)

    def update(self, G, u, v):
        for i in range(self.heapSize):
            if self.minHeap[i] == G.V[v]:
                self.minHeap[i].pi = u.value
                self.minHeap[i].key = G.W[u.value, v]
                self.decreaseKey(i, G.W[u.value, v])
                break

    def __contains__(self, a):
        return a in self.minHeap[:self.heapSize]

    def __str__(self):
        return f"{self.minHeap}"


def prim(G, root):
    for v in G.V:
        G.V[v].key = float("inf")
        G.V[v].pi = None
    G.V[root].key = 0
    heap = MinHeap(list(G.V.values()))
    A = []
    while not heap.isEmpty():
        u = heap.extractMin()
        A.append(u)
        for v in G.Adj[u.value]:
            if G.V[v] in heap and G.W[u.value, v] < G.V[v].key:
                G.V[v].pi = u.value
                G.V[v].key = G.W[u.value, v]
                heap.update(G, u, v)
    return A


def main():
    vertex = ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i']
    edges = [('a', 'b', 4), ('a', 'h', 8), ('b', 'c', 7),  ('c', 'd', 7), ('d', 'e', 9),
             ('e', 'f', 10), ('f', 'g', 2), ('g',
                                             'h', 1), ('h', 'i', 7), ('g', 'i', 6),
             ('i', 'c', 2), ('f', 'c', 4)]
    graph = initGraph(vertex, edges)
    print(prim(graph, 'a'))


if __name__ == "__main__":
    main()
