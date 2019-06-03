"""
input: canffcancd
output: can

solution: suffix array
"""


def longestRepeatedSubStr(s):
    if len(s) < 1:
        return ""
    sufStr = list(sorted([s[i:] for i in range(len(s))]))
    subStr, length, maxSubStr, maxLen = s[0], 1, s[0], 1
    for i in range(len(sufStr)-1):
        j = 0
        while j < min(len(sufStr[i]), len(sufStr[i+1])) and sufStr[i][j] == sufStr[i+1][j]:
            j += 1
        subStr, length = sufStr[i][:j], j
        if length > maxLen:
            maxSubStr, maxLen = subStr, length
    return maxSubStr


def main():
    s = "canffcancd"
    maxSubStr = longestRepeatedSubStr(s)
    print(maxSubStr)


if __name__ == "__main__":
    main()
