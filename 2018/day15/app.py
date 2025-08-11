from util.util import read_input_without_strip
import unittest

INPUT = list(filter(lambda line: line, read_input_without_strip("day15/input.txt")))


def bfs(start, targets, battlefield):
    from collections import deque

    def get_adjacent(pos):
        x, y = pos
        return [(x - 1, y), (x, y - 1), (x, y + 1), (x + 1, y)]

    queue = deque()
    visited = set()

    queue.append((start, []))
    visited.add(start)

    found_paths = []
    min_path_len = None

    while queue:
        current_pos, path = queue.popleft()

        if min_path_len is not None and len(path) > min_path_len:
            break  # Stop exploring longer paths

        if current_pos in targets:
            found_paths.append(path)
            min_path_len = len(path)
            continue

        for next_pos in get_adjacent(current_pos):
            x, y = next_pos
            if (
                0 <= x < len(battlefield)
                and 0 <= y < len(battlefield[0])
                and battlefield[x][y] == "."
                and next_pos not in visited
            ):
                visited.add(next_pos)
                queue.append((next_pos, path + [next_pos]))

    if not found_paths:
        return None

    # Sort by target position (last step), then by first step in path
    found_paths.sort(
        key=lambda path: (path[-1][0], path[-1][1], path[0][0], path[0][1])
    )
    return found_paths[0]


class Soldier:
    def __init__(self, kind, x, y, hitpoints=200, attack_power=3):
        self.kind = kind  # 'E' or 'G'
        self.hitpoints = hitpoints
        self.attack_power = attack_power
        self.position = (x, y)

    def is_enemy(self, other):
        return isinstance(other, Soldier) and self.kind != other.kind

    def is_alive(self):
        return self.hitpoints > 0

    def adjacent_positions(self):
        x, y = self.position
        return [(x - 1, y), (x, y - 1), (x, y + 1), (x + 1, y)]  # Reading order

    def has_adjacent_enemy(self, battlefield):
        return bool(self.get_adjacent_enemies(battlefield))

    def get_adjacent_enemies(self, battlefield):
        enemies = []
        for i, j in self.adjacent_positions():
            if 0 <= i < len(battlefield) and 0 <= j < len(battlefield[0]):
                other = battlefield[i][j]
                if self.is_enemy(other) and other.is_alive():
                    enemies.append(other)
        return enemies

    def move(self, battlefield):
        # Get all positions adjacent to enemies
        targets = [
            (i, j)
            for row in battlefield
            for unit in row
            if isinstance(unit, Soldier) and self.is_enemy(unit)
            for (i, j) in unit.adjacent_positions()
            if battlefield[i][j] == "."
        ]

        if not targets:
            return  # Nowhere to move

        path = bfs(self.position, targets, battlefield)
        if path:
            # Move
            next_pos = path[0]
            battlefield[next_pos[0]][next_pos[1]] = self
            battlefield[self.position[0]][self.position[1]] = "."
            self.position = next_pos

    def attack(self, battlefield):
        enemies = self.get_adjacent_enemies(battlefield)
        if not enemies:
            return

        # Choose weakest enemy (ties: reading order)
        target = min(enemies, key=lambda e: (e.hitpoints, e.position[0], e.position[1]))
        target.hitpoints -= self.attack_power
        if target.hitpoints <= 0:
            x, y = target.position
            battlefield[x][y] = "."


def parse_battlefield(inp, elf_attack_power=3):
    battlefield = []
    for i, line in enumerate(inp):
        row = []
        for j, char in enumerate(line):
            if char == "E":
                row.append(Soldier("E", i, j, attack_power=elf_attack_power))
            elif char == "G":
                row.append(Soldier("G", i, j))
            else:
                row.append(char)
        battlefield.append(row)
    return battlefield


def calculate_final_hitpoints(battlefield):
    hitpoints = 0
    for row in battlefield:
        for cell in row:
            if isinstance(cell, Soldier):
                hitpoints += cell.hitpoints
    return hitpoints


def play_round(battlefield, stop_if_elf_dies=False):
    units = []
    for row in battlefield:
        for cell in row:
            if isinstance(cell, Soldier):
                units.append(cell)

    # Sort units by reading order (position)
    units.sort(key=lambda u: (u.position[0], u.position[1]))

    for unit in units:
        if stop_if_elf_dies and unit.kind == "E" and not unit.is_alive():
            return True  # Combat ends because an elf has died

        if not unit.is_alive():
            continue

        # Check for combat end early
        enemies_remaining = any(
            isinstance(cell, Soldier) and unit.is_enemy(cell)
            for row in battlefield
            for cell in row
        )
        if not enemies_remaining:
            return True  # Combat ends during the round

        # Try to attack; if no adjacent enemies, try to move and then attack again
        if not unit.has_adjacent_enemy(battlefield):
            unit.move(battlefield)

        unit.attack(battlefield)

    return False


def count_elves(battlefield):
    count = 0
    for row in battlefield:
        for cell in row:
            if isinstance(cell, Soldier) and cell.kind == "E":
                count += 1
    return count


def part1(inp):
    battlefield = parse_battlefield(inp)
    total_rounds = 0
    while True:
        # If combat ends during the round, do not count it
        combat_ended = play_round(battlefield)
        if combat_ended:
            break
        total_rounds += 1

    return total_rounds * calculate_final_hitpoints(battlefield)


def part2(inp):
    elf_attack_power = 4
    while True:
        battlefield = parse_battlefield(inp, elf_attack_power=elf_attack_power)
        initial_elf_count = count_elves(battlefield)
        total_rounds = 0

        # Simulate combat
        while True:
            combat_ended_early = play_round(battlefield, True)
            if combat_ended_early:
                break
            total_rounds += 1

        surviving_elves = count_elves(battlefield)

        if surviving_elves == initial_elf_count:
            return total_rounds * calculate_final_hitpoints(battlefield)

        elf_attack_power += 1


class TestDay14(unittest.TestCase):
    def read_input(self, file_path):
        return list(filter(lambda line: line, read_input_without_strip(file_path)))

    def test_part1_input_1(self):
        self.assertEqual(
            36334,
            part1(self.read_input("day15/p1_input_example_1.txt")),
        )

    def test_part1_input_2(self):
        self.assertEqual(
            27755,
            part1(self.read_input("day15/p1_input_example_2.txt")),
        )

    def test_part1_input_3(self):
        self.assertEqual(
            28944,
            part1(self.read_input("day15/p1_input_example_3.txt")),
        )

    def test_part1_input_4(self):
        self.assertEqual(
            18740,
            part1(self.read_input("day15/p1_input_example_4.txt")),
        )

    def test_part2_input_1(self):
        self.assertEqual(
            4988,
            part2(self.read_input("day15/p2_input_example_1.txt")),
        )

    def test_part2_input_2(self):
        self.assertEqual(
            31284,
            part2(self.read_input("day15/p2_input_example_2.txt")),
        )

    def test_part2_input_3(self):
        self.assertEqual(
            3478,
            part2(self.read_input("day15/p2_input_example_3.txt")),
        )

    def test_part2_input_4(self):
        self.assertEqual(
            6474,
            part2(self.read_input("day15/p2_input_example_4.txt")),
        )

    def test_part2_input_5(self):
        self.assertEqual(
            1140,
            part2(self.read_input("day15/p2_input_example_5.txt")),
        )


# To run with the test inputs, run those two lines
# if __name__ == "__main__":
#     unittest.main()


# To run with the actual input, run those two lines
print("Part 1: ", part1(INPUT))
print("Part 2: ", part2(INPUT))
