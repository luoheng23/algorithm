"""
input: 
{aaa, bbb, ccc}
{bbb, ddd}
{eee, fff}
{ggg}
{ddd, hhh}
output: {aaa, bbb, ccc, ddd, hhh}, {eee, fff}, {ggg}

solution: 
"""

import copy


def mergeSets(strSets):
    cpStrSets = copy.deepcopy(strSets)
    flag = True
    while flag:
        length = len(cpStrSets)
        for i in range(length):
            for j in range(i+1, length):
                if cpStrSets[i] & cpStrSets[j]:
                    cpStrSets[i] |= cpStrSets[j]
                    del cpStrSets[j]
                    flag = False
                    break
            if not flag:
                flag = True
                break
        else:
            flag = False
    return cpStrSets


def main():
    s = [{"aaa", "bbb", "ccc"}, {"bbb", "ddd"}, {
        "eee", "fff"}, {"ggg"}, {"ddd", "hhh"}]
    newS = mergeSets(s)
    print(newS)


if __name__ == "__main__":
    main()
