
class vertex(int):
    WHITE = 0
    GRAY = 1
    BLACK = 2

    def __init__(self, value):
        self.color = vertex.WHITE
        self.d = float("inf")
        self.pi = None

class Graph:
    def __init__(self, vertex, edges):
        self.V = 