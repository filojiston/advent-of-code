from util.util import read_input

inp = read_input("day03/input.txt")
fabric_size = 1000


class Claim:
    def __init__(self, id, left, top, width, height):
        self.id = id
        self.left = left
        self.top = top
        self.width = width
        self.height = height


def part1(inp):
    claims = parse_claims(inp)
    fabric = create_fabric(claims)
    return sum(1 for row in fabric for cell in row if cell > 1)


def part2(inp):
    claims = parse_claims(inp)
    fabric = create_fabric(claims)
    try:
        return next(
            claim.id
            for claim in claims
            if all(
                fabric[i][j] == 1
                for i in range(claim.left, claim.left + claim.width)
                for j in range(claim.top, claim.top + claim.height)
            )
        )
    except StopIteration:
        return None


def parse_claims(inp):
    claims = []
    for claim in inp:
        id, _, pos, size = claim.split()
        left, top = pos[:-1].split(",")
        width, height = size.split("x")
        claims.append(Claim(id, int(left), int(top), int(width), int(height)))
    return claims


def create_fabric(claims):
    fabric = [[0 for _ in range(fabric_size)] for _ in range(fabric_size)]
    for claim in claims:
        for i in range(claim.left, claim.left + claim.width):
            for j in range(claim.top, claim.top + claim.height):
                fabric[i][j] += 1
    return fabric


print(f"Part 1: {part1(inp)}")
print(f"Part 2: {part2(inp)}")
