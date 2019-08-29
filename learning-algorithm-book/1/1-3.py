"""
print:
(a), (b), (c), (d), (e), ...... (z)
(a,b), (a,c), (a,d), ... (y,z)
...
(a,b,c,d, ...,x,y,z)
"""

import string


def calcPerm(s, temp, num, total, fro, to, array):
    if num > total:
        return
    if num == total:
        array.append(f"({','.join(temp)})")
    else:
        for i in range(fro, to+1):
            temp[num] = s[i]
            calcPerm(s, temp, num + 1, total, i + 1, to, array)


def main():
    array = []
    s = string.ascii_letters[:26]
    for i in range(1, 27):
        calcPerm(s, [None] * i, 0, i, 0, len(s)-1, array)
        print(",".join(array))
        array = []


if __name__ == "__main__":
    main()
