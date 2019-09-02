"""
input: [-2.5, 4, 0, 3, 0.5, 8, -1]
output: 3 * 0.5 * 8

solution: dynamic
"""


def MaxProductSubArray(A):
    Min, Max = A[0], A[0]
    maxResult = A[0]
    for a in A[1:]:
        end1, end2 = Min * a, Max * a
        Max = max(max(end1, end2), a)
        Min = min(min(end1, end2), a)
        maxResult = max(maxResult, Max)
    return maxResult


def main():
    A = [-2.5, 4, 0, 3, 0.5, 8, -1]
    print(MaxProductSubArray(A))


if __name__ == "__main__":
    main()
