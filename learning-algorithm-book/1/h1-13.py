"""
input: canffcancd
output: can

solution: suffix array
"""


def longestRepeatedSubStr(s):
    if len(s) < 1:
        return ""
    sufStr = list(sorted([s[i:] for i in range(len(s))]))
    maxSubStr, maxLen = s[0], 1
    for i in range(len(sufStr)-1):
        length = 0
        minLen = min(len(sufStr[i]), len(sufStr[i+1]))
        while length < minLen and sufStr[i][length] == sufStr[i+1][length]:
            length += 1
        if length > maxLen:
            maxSubStr, maxLen = sufStr[i][:length], length
    return maxSubStr


def main():
    s = "canffcancd"
    maxSubStr = longestRepeatedSubStr(s)
    print(maxSubStr)


if __name__ == "__main__":
    main()
