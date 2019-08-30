"""
input: bad, adb (alpha characters)
output: they are the same

solution: hash
1. no duplicated char:
use a bitmap as hash
2. duplicated char:
use a array as hash
"""


def no_duplicated_str(s1, s2):
    def f(x): return ord(x) - ord("A")
    has = 0
    for c in s1:
        has |= 1 << f(c)
    for c in s2:
        has ^= 1 << f(c)
    if has != 0:
        return False
    return True


def duplicated_str(s1, s2):
    def f(x): return (ord(x) - ord("A")
                      ) if "Z" >= "x" else (ord(x) - ord("a") + 26)
    s = [0 for _ in range(52)]
    for c in s1:
        s[f(c)] += 1
    for c in s2:
        s[f(c)] -= 1
    for i in s:
        if i != 0:
            return False
    return True


def main():
    s1 = "ABCDabcd"
    s2 = "aABCbcdD"
    if no_duplicated_str(s1, s2):
        print("right1")
    s1 = "aaabbbcccdddAAA"
    s2 = "AAAabcabcabcddd"
    if duplicated_str(s1, s2):
        print("right2")
    s1 = "abcde"
    s2 = "abcd"
    if not no_duplicated_str(s1, s2):
        print("right3")


if __name__ == "__main__":
    main()
