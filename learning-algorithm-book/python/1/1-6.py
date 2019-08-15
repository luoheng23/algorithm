"""
input: habbafgh
output: h, abba, f, g, h

solution: use Manacher (algorithm to find the longest string)
"""


def manacher(s, length):
    p = [0] * length
    mx = 0
    Id = 0
    for i in range(1, length):
        if mx > i:
            p[i] = min(p[2*Id-i], mx - i)
        else:
            p[i] = 1
        while i + p[i] < length and s[i+p[i]] == s[i-p[i]]:
            p[i] += 1
        if p[i] + i > mx:
            mx = p[i] + i
            Id = i
    return p


def cut_str(s, output):
    new_s = "#" + "#".join(list(s)) + "#"
    p = manacher(new_s, len(new_s))
    mx = max(p)
    Id = p.index(mx)
    if mx <= 2:
        output.extend(list(s))
    else:
        sub_str = "".join(c for c in new_s[Id-mx+1:Id+mx] if c != "#")
        new_index = s.index(sub_str)
        cut_str(s[:new_index], output)
        output.append(sub_str)
        cut_str(s[new_index+len(sub_str):], output)


def main():
    output = []
    s = "habbafgh"
    cut_str(s, output)
    print(",".join(output))


if __name__ == "__main__":
    main()
