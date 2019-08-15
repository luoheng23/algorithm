"""
input: article, {"hello", "world", "good"}
output: smallest string contains the string array

solution: always move to the right
"""


def smallestContainStr(article, strList):
    articleList = article.split(" ")
    pos = 0
    posList = [-1, -1, -1, -1]
    for word in articleList:
        index = strList.index(word)
        posList[index] = pos
        pos += len(word) + 1
        if posList[0] != -1 and posList[1] != -1 and posList[2] != -1:
            return article[min(posList[:3]):max(posList[:3])]
    return ""


def main():
    article = "this is a good world if you know say hello world that may be very good"
