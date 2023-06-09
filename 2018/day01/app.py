from util.util import read_input

frequencies = [int(x) for x in read_input("day01/input.txt")]


def part1(frequncies):
    return sum(frequencies)


def part2(frequencies):
    seen = set()
    current = 0
    while True:
        for frequency in frequencies:
            current += frequency
            if current in seen:
                return current
            seen.add(current)


print(f"Part 1: {part1(frequencies)}")
print(f"Part 2: {part2(frequencies)}")
