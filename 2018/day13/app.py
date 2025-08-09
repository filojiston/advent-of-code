from util.util import read_input_without_strip
from enum import Enum

inp = list(filter(lambda line: line, read_input_without_strip("day13/input.txt")))
inp_example = list(
    filter(lambda line: line, read_input_without_strip("day13/input_example.txt"))
)


class Direction(Enum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3

    def turn_left(self) -> "Direction":
        return Direction((self.value - 1) % 4)

    def turn_right(self) -> "Direction":
        return Direction((self.value + 1) % 4)

    def vector(self) -> tuple[int, int]:
        return {
            Direction.UP: (0, -1),
            Direction.RIGHT: (1, 0),
            Direction.DOWN: (0, 1),
            Direction.LEFT: (-1, 0),
        }[self]


class Turn(Enum):
    LEFT = 0
    STRAIGHT = 1
    RIGHT = 2

    def next(self) -> "Turn":
        order = list(Turn)
        idx = order.index(self)
        return order[(idx + 1) % len(order)]


class Cart:
    def __init__(self, x, y, direction, turn_state=Turn.LEFT):
        self.x = x
        self.y = y
        self.dir = direction
        self.turn_state = turn_state

    def __hash__(self):
        return hash(id(self))

    def __eq__(self, other):
        return self is other

    def turn_right(self):
        self.dir = self.dir.turn_right()

    def turn_left(self):
        self.dir = self.dir.turn_left()

    def update_turn_state(self):
        self.turn_state = self.turn_state.next()


def parse_input(inp):
    direction_map = {
        "^": Direction.UP,
        "v": Direction.DOWN,
        "<": Direction.LEFT,
        ">": Direction.RIGHT,
    }

    underlying_track = {
        "^": "|",
        "v": "|",
        "<": "-",
        ">": "-",
    }

    grid = []
    carts = []

    for y, line in enumerate(inp):
        row = []
        for x, char in enumerate(line):
            if char in direction_map:
                carts.append(Cart(x, y, direction_map[char]))
                char = underlying_track[char]
            row.append(char)
        grid.append(row)
    return grid, carts


def move_cart(cart, grid):
    dx, dy = cart.dir.vector()
    cart.x += dx
    cart.y += dy

    tile = grid[cart.y][cart.x]

    if tile == "/":
        if cart.dir in (Direction.UP, Direction.DOWN):
            cart.turn_right()
        else:
            cart.turn_left()

    elif tile == "\\":
        if cart.dir in (Direction.UP, Direction.DOWN):
            cart.turn_left()
        else:
            cart.turn_right()

    elif tile == "+":
        if cart.turn_state == Turn.LEFT:
            cart.turn_left()
        elif cart.turn_state == Turn.RIGHT:
            cart.turn_right()

        cart.update_turn_state()


def tick(grid, carts, stop_on_first_collision=True):
    carts.sort(key=lambda c: (c.y, c.x))

    positions = {(c.x, c.y): c for c in carts}
    carts_to_remove = set()

    for cart in carts:
        if cart in carts_to_remove:
            continue

        positions.pop((cart.x, cart.y))

        move_cart(cart, grid)
        pos = (cart.x, cart.y)

        if pos in positions:
            other_cart = positions[pos]

            carts_to_remove.add(cart)
            carts_to_remove.add(other_cart)
            positions.pop(pos)

            if stop_on_first_collision:
                return (pos, carts)
        else:
            positions[pos] = cart

    carts = [c for c in carts if c not in carts_to_remove]

    if not stop_on_first_collision:
        if len(carts) == 1:
            last_cart = carts[0]
            return ((last_cart.x, last_cart.y), carts)
        return (None, carts)

    return (None, None)


def part1(inp):
    grid, carts = parse_input(inp)
    while True:
        collision_pos, _ = tick(grid, carts, stop_on_first_collision=True)
        if collision_pos:
            return collision_pos


def part2(inp):
    grid, carts = parse_input(inp)
    while True:
        last_cart_pos, updated_carts = tick(grid, carts, stop_on_first_collision=False)
        if last_cart_pos:
            return last_cart_pos
        else:
            carts = updated_carts


print(f"Part 1: {part1(inp)}")
print(f"Part 2: {part2(inp)}")
