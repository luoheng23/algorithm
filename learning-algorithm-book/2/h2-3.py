"""
input: numbers
output: 2 appear once number

solution: ^
"""

from functools import reduce
import random 

findTwoOnceNum = __import__("h2-2").findTwoOnceNum


def main():
    s = list(range(9)) + list(range(9)) + list(range(10, 12))
    random.shuffle(s)
    print(findTwoOnceNum(reduce(lambda x, y: x ^ y, s), s))


if __name__ == "__main__":
    main()
