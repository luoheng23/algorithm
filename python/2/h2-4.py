"""
input: n
output times of 1 in 1..n

solution:
1..9: 1
1..99: 9 * 1 + 10
1..999: 9 * (9 * 1 + 10) + 100
1..99.99(n): 9 * (n - 1) + 100.00(n)
"""

def countOneTimes(n):
    lenN = len(str(n))
    cntN = [1 for _ in range(lenN)]
    for i in range(1, lenN):
        cntN[i] = cntN[i-1] * 9 + 10**i
    sumN = 0
    for i, c in enumerate(str(n)):
        at = lenN - i
        if c > '1':
            pass


