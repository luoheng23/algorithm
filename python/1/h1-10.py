"""
input: {5, 10, 20, 15, 25, 30}, {5, 15, 35, 25}
output: {10, 20, 30}

solution: use intersection set
"""


def intersectionSet(A, B):
    A = A - B
    return A


def main():
    A = {5, 10, 20, 15, 25, 30}
    B = {5, 15, 35, 25}
    newA = intersectionSet(A, B)
    print(newA)


if __name__ == "__main__":
    main()
