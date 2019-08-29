"""
input: xxxyyyyyyz
output: 3x6yz

solution: O(n)
"""


def zipString(s):
    # make the last character same as others
    new_s = s + "$"
    length = len(new_s)
    if length == 0:
        return ""
    curCha, numCha = new_s[0], 1
    zipStrList = []
    for c in new_s[1:]:
        if c == curCha:
            numCha += 1
        else:
            if numCha != 1:
                zipStrList.append(str(numCha))
            zipStrList.append(curCha)
            curCha, numCha = c, 1
    return "".join(zipStrList)


def main():
    s = "xxxyyyyyyz"
    new_s = zipString(s)
    print(new_s)


if __name__ == "__main__":
    main()
