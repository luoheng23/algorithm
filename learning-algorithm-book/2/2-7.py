"""
input: RGBRBGGBR
output: RRRGGGBBB

solution: quick sort
"""


def partition(s):
    length = len(s)
    sta, cur, end = 0, 0, length - 1
    while cur < end:
        if s[cur] == "R":
            s[cur], s[sta] = s[sta], s[cur]
            sta += 1
            cur += 1
        elif s[cur] == "B":
            s[cur], s[end] = s[end], s[cur]
            end -= 1
        else:
            cur += 1


def main():
    s = list("RGBRBGGBR")
    partition(s)
    print(s)


if __name__ == "__main__":
    main()
