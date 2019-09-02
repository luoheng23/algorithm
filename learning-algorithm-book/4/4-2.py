"""
input: sorted Matrix, key
output: key in Matrix ?

solution: search from right-up position
"""

def find(Matrix, key) -> bool:
    width, height = len(Matrix[0]), len(Matrix)
    startX = 0
    startY = width - 1
    while 0 <= startX < height and 0 <= startY < width and Matrix[startX][startY] != key:
        if Matrix[startX][startY] > key:
            startY -= 1
        else:
            startX += 1
    if not (0 <= startX < height and 0 <= startY < width):
        return False
    return True

def main():
    matrix = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]
    if find(matrix, 7):
        print("7 is in matrix")
    if not find(matrix, 10):
        print("10 is not in matrix")

if __name__ == "__main__":
    main()