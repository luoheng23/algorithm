"""
input: abaccdeff
output: b

solution: find the first character that only appears once
"""

def FirstChar(s: str) -> str:
    def f(x): return 26+ord(x)-ord('a') if 'a' <= x <= 'z' else ord(x) - ord('A')
    L = [0] * 52
    for c in s:
        L[f(c)] += 1
    for c in s:
        if L[f(c)] == 1:
            return c

def main():
    s = "abaccdeff"
    print(FirstChar(s))

if __name__ == "__main__":
    main()