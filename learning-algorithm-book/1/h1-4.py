"""
input: "J* Smi??"
output: "John Smith"
"""


def oneMatch(pattern, string, start, len_str, len_pat):
    i = 0
    while i < len_pat and start < len_str:
        if pattern[i] == "*" and findInArticle(pattern[i+1:], string[start:]):
            return True
        elif pattern[i] == "?" or pattern[i] == string[start]:
            i += 1
            start += 1
        else:
            break
    if i == len(pattern):
        return True
    return False


def findInArticle(pattern, article):
    for i in range(len(article)):
        if oneMatch(pattern, article, i, len(article), len(pattern)):
            return True
    return False


def main():
    pattern = "J* Smi??"
    s = "John Smithsafsdfsa fsadfsdafasfJlsdfjksldf jflsdjfksjflaskjflsfsdflajsfljSmifksd;jfsklfa;jaslfksfjlasdjflasfj"
    pos = findInArticle(pattern, s)
    print(pos)


if __name__ == "__main__":
    main()
