"""
input: "J* Smi??"
output: "John Smith"
"""


def oneMatch(pattern, string, len_p, len_s):
    p_start, s_start = 0, 0
    while p_start < len_p and s_start < len_s:
        if pattern[p_start] == "*":
            i, len_i = findInArticle(pattern[p_start+1:], string[s_start:], len_p-p_start-1, len_s-s_start)
            if i != -1:
                return s_start + i + len_i
        elif pattern[p_start] == string[s_start] or pattern[p_start] == "?" and string[s_start] != " ":
            p_start += 1
            s_start += 1
        else:
            break
    if p_start == len(pattern):
        return s_start
    return 0


def findInArticle(pattern, article, len_p, len_s):
    for i in range(len(article)):
        len_i = oneMatch(pattern, article[i:], len_p, len_s - i)
        if len_i != 0:
            return i, len_i
    return -1, 0


def main():
    p = "J* Smi??"
    s = " safsdfJohnsa fsadfsdafa SmithsfJlsdfjksldf jflsdjfksjflaskjflsfsdflajsfljSmifksd;jfsklfa;jaslfksfjlasdjflasfj"
    i, len_i = findInArticle(p, s, len(p), len(s))
    print(s[i:i+len_i], i, len_i)


if __name__ == "__main__":
    main()
