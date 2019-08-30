"""
move * to the left, while other characters don't change position

solution: read the string from the bottom
"""


def move(s):
    index = len(s) - 1
    for i in range(len(s)-1, -1, -1):
        if s[i] != "*":
            s[i], s[index] = s[index], s[i]
            index -= 1


def main():
    s = list("**a*b*c*defg***h**i***jklm***")
    move(s)
    print("".join(s))


if __name__ == "__main__":
    main()
