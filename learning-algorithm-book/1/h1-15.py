"""
input: 010111
output: 01, 011, 1

solution: use n/2 window
"""

def cut01(s: str):
    num = [0, 0]
    for c in s:
        num[ord(c) - ord('0')] += 1
    half = sum(num) // 2
    half0 = num[0] / 2
    cutNum = 0 # only count 0
    for i in range(half):
        if s[i] == '0':
            cutNum += 1
    if cutNum == half0:
        return [half]
    for i in range(half, sum(num)):
        if s[i-half] == '0':
            cutNum -= 1
        if s[i] == '0':
            cutNum += 1
        if cutNum == half0:
            return [i - half + 1, i + 1]

def main():
    s = "010111"
    cut = cut01(s)
    if len(cut) == 1:
        strs = [s[:cut[0]], s[cut[0]:]]
    else:
        strs = [s[:cut[0]], s[cut[0]:cut[1]], s[cut[1]:]]
    print(strs)

if __name__ == "__main__":
    main()
