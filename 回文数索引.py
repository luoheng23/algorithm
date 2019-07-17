N = int(input())
for _ in range(N):
    s = input()
    ls = len(s)
    if ls <= 0:
        print(-1)
    p = -1
    for i in range(ls//2+1):
        if s[i] != s[-i-1]:
            if s[i+1] == s[-i-1]:
                p = i
            else:
                p = ls - i - 1
            break
    print(p)