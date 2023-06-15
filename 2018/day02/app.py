from util.util import read_input
from collections import Counter

box_ids = read_input("day02/input.txt")


def part1(ids):
    exactly_two = 0
    exactly_three = 0
    for box_id in ids:
        counts = Counter(box_id)
        if 2 in counts.values():
            exactly_two += 1
        if 3 in counts.values():
            exactly_three += 1
    return exactly_two * exactly_three


def part2(ids):
    for box_id in ids:
        for other_id in ids:
            if box_id == other_id:
                continue
            diff = sum(ch1 != ch2 for ch1, ch2 in zip(box_id, other_id))
            if diff == 1:
                return "".join(ch1 for ch1, ch2 in zip(box_id, other_id) if ch1 == ch2)
    return None


print(f"Part 1: {part1(box_ids)}")
print(f"Part 2: {part2(box_ids)}")
