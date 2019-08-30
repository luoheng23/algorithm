"""
input: They are students, aeiou
output: Thy r stdnts.

solution: use bit (dict)
"""


def trimStr(s, trimS):
    def f(x): return ord(x) - ord("A") + \
        1 if "A" <= x <= "Z" else ord(
            x) - ord("a") + 27 if "a" <= x <= "z" else 0
    bit = 0
    for c in trimS:
        bit |= (1 << f(c))
    newStrList = [c for c in s if not (bit & (1 << f(c)))]
    return "".join(newStrList)


def main():
    s = "They are students."
    trimS = "aeiou"
    newS = trimStr(s, trimS)
    print(newS)


if __name__ == "__main__":
    main()
