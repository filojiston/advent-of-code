from util.util import read_input
from datetime import datetime

inp = read_input("day04/input.txt")


def part1(inp):
    shifts = parse_shifts(inp)
    guards = calculate_sleep_map(shifts)

    guard = max(guards, key=lambda x: sum(guards[x]))
    minute = guards[guard].index(max(guards[guard]))
    return int(guard) * minute


def part2(inp):
    shifts = parse_shifts(inp)
    guards = calculate_sleep_map(shifts)

    guard = max(guards, key=lambda x: max(guards[x]))
    minute = guards[guard].index(max(guards[guard]))
    return int(guard) * minute


def parse_shifts(inp):
    return sorted(inp, key=lambda x: datetime.strptime(x[1:17], "%Y-%m-%d %H:%M"))


def calculate_sleep_map(shifts):
    guards = {}
    guard = None
    for shift in shifts:
        if "Guard" in shift:
            guard = shift.split(" ")[3][1:]
            if guard not in guards:
                guards[guard] = [0] * 60
        elif "falls asleep" in shift:
            start = int(shift[15:17])
        elif "wakes up" in shift:
            end = int(shift[15:17])
            for i in range(start, end):
                guards[guard][i] += 1
    return guards


print("Part 1: ", part1(inp))
print("Part 2: ", part2(inp))
