
import sys
sys.path.insert(0, "..")
from chapter22.graph import initGraph


def init(G, s):
    for v in G.V:
        G.V[v].d = float("inf")
        G.V[v].pi = None
    G.V[s].d = 0


def relax(G, u, v):
    if G.V[v].d > G.V[u].d + G.W[u, v]:
        G.V[v].d = G.V[u].d + G.W[u, v]
        G.V[v].pi = u
