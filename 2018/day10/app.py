from util.util import read_input

import re
from dataclasses import dataclass


@dataclass
class Point:
    pos: tuple[int, int]
    vel: tuple[int, int]

    def move(self):
        self.pos = (self.pos[0] + self.vel[0], self.pos[1] + self.vel[1])


def parse_input(inp):
    regex = r"position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>"
    points = []
    for line in inp:
        match = re.search(regex, line)
        if match:
            x, y, velx, vely = map(int, match.groups())
            points.append(Point((x, y), (velx, vely)))
    return points


def get_area(points):
    xs = [p.pos[0] for p in points]
    ys = [p.pos[1] for p in points]
    return (max(xs) - min(xs)) * (max(ys) - min(ys))


def print_points(points):
    xs = [p.pos[0] for p in points]
    ys = [p.pos[1] for p in points]
    min_x, max_x = min(xs), max(xs)
    min_y, max_y = min(ys), max(ys)

    if max_x - min_x > 100 or max_y - min_y > 100:
        print("Area too large to display meaningfully")
        return

    point_set = {p.pos for p in points}

    for y in range(min_y, max_y + 1):
        line = ""
        for x in range(min_x, max_x + 1):
            line += "#" if (x, y) in point_set else "."
        print(line)


def find_message(points, max_steps=15000):
    min_area = float("inf")
    min_step = 0
    min_points = None

    current_points = [Point(p.pos, p.vel) for p in points]

    for step in range(max_steps):
        current_area = get_area(current_points)

        if current_area < min_area:
            min_area = current_area
            min_step = step
            min_points = [Point(p.pos, p.vel) for p in current_points]

        for point in current_points:
            point.move()

        if current_area < min_area * 0.9:
            min_area = current_area
        elif current_area > min_area * 1.5:
            break

    return min_step, min_points


def solve(inp):
    points = parse_input(inp)
    step, message_points = find_message(points)
    print(f"\nMessage appears at step {step}:")
    print_points(message_points)
    return step


inp = read_input("day10/input.txt")
solve(inp)
