"""
input: string
output: string times

solution: use an array to count times
"""


def count(s):
    def f(x): return ord(x) - ord("A") if x <= "Z" else ord(x) - ord("a") + 26

    def back_f(x): return chr(
        x + ord("A")) if x <= 25 else chr(x + ord("a") - 26)

    array = [0] * 52
    for c in s:
        array[f(c)] += 1

    return [(back_f(i), c) for i, c in enumerate(array) if c != 0]


def main():
    s = "AJFDFJKfjfjdksfjskdjffjafhafJFkkdfjsdjfsdlfjsdfjhwofiuwerejfldsf"
    print(count(s))


if __name__ == "__main__":
    main()
