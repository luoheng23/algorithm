"""
input: wepicabc, epiabcni
output: epi

solution: suffix array
"""


def maxIntersectSubStr(A, B):
    sufStrA = [(A[i:], 0) for i in range(len(A))]
    sufStrB = [(B[i:], 1) for i in range(len(B))]
    newSufStr = sufStrA + sufStrB
    newSufStr.sort(key=lambda x: x[0])
    maxSubStr, maxLen, pos = "", 0, -1
    for i in range(len(newSufStr)-1):
        # should not be the same array
        if newSufStr[i][1] != newSufStr[i+1][1]:
            length = 0
            minLen = min(len(newSufStr[i][0]), len(newSufStr[i+1][0]))
            while length < minLen and newSufStr[i][0][length] == newSufStr[i+1][0][length]:
                length += 1
            if length > maxLen or length == maxLen and minLen > pos:
                maxSubStr, maxLen, pos = newSufStr[i][0][:length], length, minLen
    return maxSubStr


def main():
    A, B = "wepicabc", "epiabcni"
    maxSubStr = maxIntersectSubStr(A, B)
    print(maxSubStr)


if __name__ == "__main__":
    main()
