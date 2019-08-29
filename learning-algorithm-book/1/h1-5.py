"""
input: abc     efg    hij
output: cba gfe jih

solution: first reverse the string, then move it to the specific position
"""

# this is a function to reverse string in place, it's from 1-1.py
reverse = __import__("1-1").reverse


def zipAndReverse(s):
    length = len(s)
    start, paste = 0, 0
    flag = True
    for i in range(length):
        if flag and s[i] == " ":
            reverse(s, start, i-1)
            # move to the right place
            if start != paste:
                for j in range(start, i):
                    s[paste] = s[j]
                    paste += 1
                    j += 1
                paste += 1
            else:
                paste = i + 1
            flag = False
        elif s[i] != " " and not flag:
            flag = True
            start = i
    reverse(s, start, length-1)
    if start != paste:
        for j in range(start, length):
            s[paste] = s[j]
            paste += 1
            j += 1
    else:
        paste = length
    return s[:paste]


def main():
    s = list("abc     efg      hij")
    newS = zipAndReverse(s)
    print("".join(newS))


if __name__ == "__main__":
    main()
