"""
input: a1, a2, ..., an, (i, j, k)
output: ai, ai+1, ..., aj k min numbers

solution: quick sort
"""


def partition(s, left, right):
    if right - left <= 1:
        return left
    i, j, pivot = left, right-1, s[(left+right)//2]
    while i < j:
        if s[i] <= pivot:
            i += 1
        else:
            s[i], s[j] = s[j], s[i]
            j -= 1
    return i


def quickSelect(s, k, left, right):
    if right > left:
        p = partition(s, left, right)
        if p == k:
            return s[:p]
        elif p < k:
            return quickSelect(s, k, p, right)
        else:
            return quickSelect(s, k, left, p)


def smallestKNum(s, k):
    return quickSelect(s, k, 0, len(s))


def main():
    s = [1, 4, 2, 5, 6, 3, 9, 7, 8, 10]
    print(smallestKNum(s, 4))


if __name__ == "__main__":
    main()
