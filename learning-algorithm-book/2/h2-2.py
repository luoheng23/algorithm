"""
input: numbers
output: 3 number appear once 

solution: ^
"""

import random
from functools import reduce


def findTheLastOne(x):
    return x & (-x)

def findTwoOnceNum(x, s):
    num = findTheLastOne(x)
    a, b = 0, 0
    for i in s:
        if i & num:
            a ^= i
        else:
            b ^= i
    return a, b


def findThreeOnceNum(x, s):
    num = findTheLastOne(x)
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
    print(s, findThreeOnceNum(reduce(lambda x, y: x ^ y, s), s))


if __name__ == "__main__":
    main()
