"""
input: article, {"hello", "world", "good"}
output: smallest string contains the string array

solution: always move to the right
"""


def smallestContainStr(article, strList):
    articleList = article.split()
    pos = 0
    posList = [-1, -1, -1]
    for word in articleList:
        pos += len(word) + 1
        try:
            index = strList.index(word)
        except ValueError:
            continue
        posList[index] = pos - len(word) - 1
        if posList[0] != -1 and posList[1] != -1 and posList[2] != -1:
            return article[min(posList):max(posList)+len(word)+1]
    return ""


def main():
    article = "this is a good world if you know say hello world good may be very good"
    strList = ["world", "hello", "good"]
    print(smallestContainStr(article, strList))

if __name__ == "__main__":
    main()