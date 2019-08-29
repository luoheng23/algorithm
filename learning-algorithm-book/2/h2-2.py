"""
input: numbers
output: 3 once appear once 

solution: ^
"""

import random
from functools import reduce


def findTheLastOne(x):
    m = 0
    if x == 0:
        return 0
    while True:
        if x & (1 << m):
            return m
        m += 1


def findTwoOnceNum(x, s):
    m = findTheLastOne(x)
    num = 1 << m
    a, b = 0, 0
    for i in s:
        if i & num:
            a ^= i
        else:
            b ^= i
    return a, b


def findThreeOnceNum(x, s):
    m = findTheLastOne(x)
    num = 1 << m
    c = 0
    for i in s:
        if i & num:
            c ^= i
    newX = x ^ c
    a, b = findTwoOnceNum(newX, s + [c])
    return a, b, c


def main():
    s = list(range(9)) + list(range(9)) + list(range(10, 13))
    random.shuffle(s)
    print(findThreeOnceNum(reduce(lambda x, y: x ^ y, s), s))


if __name__ == "__main__":
    main()
