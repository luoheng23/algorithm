"""
input: aabcc, dbbca -> aadbbcbcac
output: true

solution: dynamic
"""


def IsInterleave(s1, s2, s3):
    if len(s1) + len(s2) != len(s3):
        return False
    dp = [[False] * (len(s2) + 1) for _ in range(len(s1)+1)]
    for i in range(len(s1) + 1):
        for j in range(len(s2) + 1):
            if (i - 1 >= 0 and dp[i-1][j] and s1[i-1] == s3[i+j-1]) or \
                    (j - 1 >= 0 and dp[i][j-1] and s2[j-1] == s3[i+j-1]):
                dp[i][j] = True
    return dp[len(s1)][len(s2)]

def main():
    s1, s2, s3 = "aabcc", "dbbca", "aadbbcbcac"
    if IsInterleave(s1, s2, s3):
        print("s1, s2, s3 is right")
    if not IsInterleave(s1.replace("a", "c"), s2, s3):
        print("ok")

if __name__ == "__main__":
    main()