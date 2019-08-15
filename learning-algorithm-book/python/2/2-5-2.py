"""
input: 100, 1, 2, 5, 10
output: 100元零钱的组合数

solution: 动态规划
"""


def numOfS():
    s = 100
    num = [0] * 101
    num[1], num[2] = 1, 2
    for i in range(3, 101):
        num[i] = num[i-1] + num[i-2] + num[i-5] + num[i-10]
    return num[100]


def main():
    print(numOfS())


if __name__ == "__main__":
    main()
