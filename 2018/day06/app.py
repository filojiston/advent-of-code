from util.util import read_input
from collections import defaultdict

coords = list(map(lambda x: tuple(map(int, x.split(","))), read_input("day06/input.txt")))

def manhattan_distance(p1, p2):
    return abs(p1[0] - p2[0]) + abs(p1[1] - p2[1])

def part1(coords):
    min_x = min(point[0] for point in coords)
    max_x = max(point[0] for point in coords)
    min_y = min(point[1] for point in coords)
    max_y = max(point[1] for point in coords)

    area_counts = defaultdict(int)
    infinite_regions = set()

    for x in range(min_x, max_x + 1):
        for y in range(min_y, max_y + 1):
            distances = [(manhattan_distance((x, y), point), i) for i, point in enumerate(coords)]
            distances.sort()

            if distances[0][0] < distances[1][0]:
                area_counts[distances[0][1]] += 1

                if x == min_x or x == max_x or y == min_y or y == max_y:
                    infinite_regions.add(distances[0][1])

    largest_finite_area = max(area_counts[i] for i in range(len(coords)) if i not in infinite_regions)
    return largest_finite_area

def part2(coords, threshold=10000):
    min_x = min(point[0] for point in coords)
    max_x = max(point[0] for point in coords)
    min_y = min(point[1] for point in coords)
    max_y = max(point[1] for point in coords)

    safe_region_size = 0

    for x in range(min_x, max_x + 1):
        for y in range(min_y, max_y + 1):
            total_distance = sum(manhattan_distance((x, y), point) for point in coords)
            if total_distance < threshold:
                safe_region_size += 1

    return safe_region_size

print("Part 1: ", part1(coords))
print("Part 2: ", part2(coords))
