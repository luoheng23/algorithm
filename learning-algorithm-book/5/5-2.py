"""
input: Julytype, Jultypt
output: 2

solution: dynamic
"""


def editDistance(s1, s2):
    d = [[0] * len(s2) for _ in range(len(s1))]
    for i in range(len(s1)):
        d[i][0] = i + 1
    for j in range(len(s2)):
        d[0][j] = j + 1
    for i in range(len(s1)):
        for j in range(len(s2)):
            if s1[i] == s2[j]:
                d[i][j] = d[i-1][j-1]
            else:
                d[i][j] = 1 + min(d[i][j-1], d[i-1][j], d[i-1][j-1])
    return d[len(s1)-1][len(s2)-1]


def main():
    s1 = "Julytype"
    s2 = "Jultypt"
    print(editDistance(s1, s2))


if __name__ == "__main__":
    main()
