"""
input: Array[n], S
output: shortest sub array, sum >= S

solution: iteration from 0 to n
"""


def shortestSubArray(A, S):
    s = 0
    max_s, max_e = 0, len(A)
    Sum = 0
    for e, a in enumerate(A):
        Sum += a
        while Sum >= S:
            if e - s < max_e - max_s:
                max_s, max_e = s, e
            Sum -= A[s]
            s += 1
    return max_s, max_e + 1


def main():
    A = [12, 3, 4, 5, 6, 7, 13, 2, 7, 8, 9]
    S = 21
    s, e = shortestSubArray(A, S)
    print(A[s:e])


if __name__ == "__main__":
    main()
