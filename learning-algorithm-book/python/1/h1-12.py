"""
input: aaaabbcc
output: 4

solution: O(n)
"""


def longestStr(string):
    if len(string) < 1:
        return 0
    curCha, numCha, maxCha = string[0], 1, 1
    for c in string[1:]:
        if c == curCha:
            numCha += 1
        else:
            maxCha = max(maxCha, numCha)
            numCha, curCha = 1, c
    return maxCha


def main():
    s1 = "aaaabbcc"
    s2 = "aabb"
    s3 = "ab"
    print(longestStr(s1), longestStr(s2), longestStr(s3))


if __name__ == "__main__":
    main()
