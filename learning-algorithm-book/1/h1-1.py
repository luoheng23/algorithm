"""
input: July
output yluJ
"""


def reverse(s, start, end):
    while start < end:
        s[start], s[end] = s[end], s[start]
        start += 1
        end -= 1


def main():
    s = list("July")
    reverse(s, 0, len(s)-1)
    print("".join(s))


if __name__ == "__main__":
    main()
