"""
input: madam
output: true
input: maddam
output: true
input: maaaamm
output: false
solution:
start to end
"""


def isPalin(s: str) -> bool:
    lenS = len(s)
    if lenS <= 1:
        return True
    for i, j in enumerate(s):
        if j != s[lenS-i-1]:
            return False
    return True

def main():
    s = "madam"
    s2 = "maddam"
    s3 = "maaaamm"
    print(isPalin(s), isPalin(s2), isPalin(s3))


if __name__ == "__main__":
    main()
