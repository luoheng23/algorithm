"""
convert a string to a number

solution:
1. sign
2. n = n * 10 + c
"""

def convert(s):
    n = 0
    start = 1
    if s[0] == "+":
        sign = 1
    if s[0] == "-":
        sign = -1
    else:
        sign = 1
        start = 0
    for i in range(start, len(s)):
        n = n * 10 + (ord(s[i]) - ord('0'))
    return n if sign > 0 else -n

def main():
    s = '-12345689'
    n = convert(s)
    print(n)

if __name__ == "__main__":
    main()