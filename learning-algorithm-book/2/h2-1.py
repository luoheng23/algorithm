"""
input: 1-1000 put in 1001 array
output: the only repeated number

solution: sort
"""

import random


def findRepeatedNum(s):
    while s[s[0]] != s[0]:
        s[s[0]], s[0] = s[0], s[s[0]]
    return s[0]


def main():
    s = list(range(1, 1001))
    repeatNum = random.randint(1, 1000)
    s.append(repeatNum)
    random.shuffle(s)
    num = findRepeatedNum(s)
    print(num, repeatNum)


if __name__ == "__main__":
    main()
