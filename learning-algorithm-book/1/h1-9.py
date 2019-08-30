"""
input: 
{aaa, bbb, ccc}
{bbb, ddd}
{eee, fff}
{ggg}
{ddd, hhh}
output: {aaa, bbb, ccc, ddd, hhh}, {eee, fff}, {ggg}

solution: 
"""

import copy


def makeSet(x):
    x.p = x
    x.rank = 0


def Union(x, y):
    link(FindSet(x), FindSet(y))


def link(x, y):
    if x.rank > y.rank:
        y.p = x
    else:
        x.p = y
        if x.rank == y.rank:
            y.rank += 1


def FindSet(x):
    if x != x.p:
        x.p = FindSet(x.p)
    return x.p


def mergeSets(eles, sets):
    for ele in eles:
        makeSet(ele)
    for 


def main():
    ele = ["", "aaa", "bbb", "ccc", "ddd", "eee", "fff", "ggg", "hhh"]
    
    sets = [(1, 2, 3), (2, 4), (5, 6), (7, ), (4, 8)]
    newS = mergeSets(eles, sets)
    print(newS)


if __name__ == "__main__":
    main()
