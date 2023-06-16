from util.util import read_input
import re

inp = read_input("day05/input.txt")[0]


def part1(inp):
    return react(inp)


def part2(inp):
    unique_chars = set(inp.lower())
    return min([react(re.sub(f"[{c}{c.upper()}]", "", inp)) for c in unique_chars])


def react(polymer):
    result = ""
    for c in polymer:
        if result and c != result[-1] and c.lower() == result[-1].lower():
            result = result[:-1]
        else:
            result += c
    return len(result)


print("Part 1: ", part1(inp))
print("Part 2: ", part2(inp))
