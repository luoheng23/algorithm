"""
input: [3, -4, 2, 4, -3, -1, 7, -6, 5]
output: 30240

description: don't use divide, biggest (n - 1) multiplies 

solution: dynamic
"""


def maxMultiply(A):
    s1, s2 = [1] * len(A), [1] * len(A)
    for i in range(1, len(A)):
        s1[i] *= s1[i-1] * A[i-1]
    for i in range(len(A)-2, -1, -1):
        s2[i] *= s2[i+1] * A[i+1]
    for i in range(len(s1)):
        s1[i] *= s2[i]
    return max(s1)


def main():
    A = [3, -4, 2, 4, -3, -1, 7, -6, 5]
    # from functools import reduce
    # print(reduce(lambda x, y: x * y, A))
    print(maxMultiply(A))


if __name__ == "__main__":
    main()
