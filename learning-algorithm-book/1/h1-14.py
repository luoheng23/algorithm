"""
input: wepiabc, pabcni
output: abc

solution: suffix array
"""


def maxIntersectSubStr(A, B):
    sufStrA = [(A[i:], 0) for i in range(len(A))]
    sufStrB = [(B[i:], 1) for i in range(len(B))]
    newSufStr = sufStrA + sufStrB
    newSufStr.sort(key=lambda x: x[0])
    subStr, length, maxSubStr, maxLen = "", 0, "", 0
    for i in range(len(newSufStr)-1):
        if newSufStr[i][1] != newSufStr[i+1][1]:
            j = 0
            while j < min(len(newSufStr[i][0]), len(newSufStr[i+1][0])) and newSufStr[i][0][j] == newSufStr[i+1][0][j]:
                j += 1
            subStr, length = newSufStr[i][0][:j], j
            if length > maxLen:
                maxSubStr, maxLen = subStr, length
    return maxSubStr


def main():
    A, B = "wepiabc", "pabcni"
    maxSubStr = maxIntersectSubStr(A, B)
    print(maxSubStr)


if __name__ == "__main__":
    main()
