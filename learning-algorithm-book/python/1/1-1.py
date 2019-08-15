"""
input: I am a student.
output: student. a am I

solution:
1. reverse all words
2. reverse the whole sentence
"""


def reverse(s, start, end):
    while start < end:
        s[start], s[end] = s[end], s[start]
        start += 1
        end -= 1


def rotate(s, length):
    start = 0
    for i in range(length):
        if s[i] == " ":
            reverse(s, start, i-1)
            start = i + 1
        if i == length - 1:
            reverse(s, start, i)
    reverse(s, 0, length-1)


def main():
    s = list("I am a student.")
    rotate(s, len(s))
    print("".join(s))


if __name__ == "__main__":
    main()
