"""
input: baca
output: 

description:
a - y 25
1 - 4
a, aa, aaa, ...
0, 1,  2,   ...

solution: Count
"""

def decode(n: int):



def encode(s: str):
    length = 4
    n = 1
    for c in s:
        n += (ord(c) - ord('a')) * 25 ** length
        length -= 1
    return n

def main():
    s = "baca"
    des = 12345
    print(encode(s))
    print(decode(encode(s)))
    print(decode(des))

if __name__ == "__main__":
    main()